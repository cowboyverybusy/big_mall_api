package container

import (
	"gorm.io/gorm"
	"sync"
)

// Model 通用模型接口
type Model interface {
	TableName() string // 获取表名
	DbName() string    // 获取数据库配置名称
}

// GormRepository 通用CRUD仓库(泛型实现)
type GormRepository[M Model] struct {
	db    *gorm.DB
	model M
}

// NewGormRepository 创建新仓库实例
func NewGormRepository[M Model](db *gorm.DB, model M) *GormRepository[M] {
	return &GormRepository[M]{
		db:    db,
		model: model,
	}
}

// 以下是CRUD方法实现
func (r *GormRepository[M]) Create(entity *M) error {
	return r.db.Create(entity).Error
}

func (r *GormRepository[M]) Update(entity *M) error {
	return r.db.Save(entity).Error
}

func (r *GormRepository[M]) Delete(id uint) error {
	var entity M
	return r.db.Delete(&entity, id).Error
}

func (r *GormRepository[M]) GetByID(id uint) (*M, error) {
	var entity M
	err := r.db.First(&entity, id).Error
	return &entity, err
}

func (r *GormRepository[M]) List() ([]*M, error) {
	var entities []*M
	err := r.db.Find(&entities).Error
	return entities, err
}

func (r *GormRepository[M]) Where(query interface{}, args ...interface{}) ([]*M, error) {
	var entities []*M
	err := r.db.Where(query, args...).Find(&entities).Error
	return entities, err
}

// RepositoryContainer 仓库容器(管理多个数据库的仓库)
type RepositoryContainer struct {
	mu           sync.RWMutex
	dbMap        map[string]*gorm.DB               // 数据库连接映射
	repositories map[string]map[string]interface{} // 分级存储: dbName -> modelName -> repository
	//repositories map[string]map[string]*GormRepository // 分级存储: dbName -> modelName -> repository
}

// NewRepositoryContainer 创建新容器
func NewRepositoryContainer(dbs map[string]*gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		dbMap:        dbs,
		repositories: make(map[string]map[string]interface{}),
	}
}

// Register 注册模型到容器
func (c *RepositoryContainer) Register(model Model) {
	c.mu.Lock()
	defer c.mu.Unlock()

	dbName := model.DbName()
	modelName := model.TableName()

	// 初始化数据库层级
	if _, ok := c.repositories[dbName]; !ok {
		c.repositories[dbName] = make(map[string]interface{})
	}

	// 如果已注册则跳过
	if _, exists := c.repositories[dbName][modelName]; exists {
		return
	}

	// 创建新仓库
	if db, ok := c.dbMap[dbName]; ok {
		repo := NewGormRepository(db, model)
		c.repositories[dbName][modelName] = repo
	}
}

// getRaw 获取原始仓库对象（内部使用）
func (c *RepositoryContainer) getRaw(dbName, modelName string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if dbRepos, ok := c.repositories[dbName]; ok {
		if repo, exists := dbRepos[modelName]; exists {
			return repo, true
		}
	}
	return nil, false
}

// GetDB 获取数据库连接
func (c *RepositoryContainer) GetDB(dbName string) (*gorm.DB, bool) {
	db, ok := c.dbMap[dbName]
	return db, ok
}

// GetRepository 获取模型仓库（泛型函数）
func GetRepository[M Model](container *RepositoryContainer, model M) (*GormRepository[M], bool) {
	dbName := model.DbName()
	modelName := model.TableName()

	if repo, ok := container.getRaw(dbName, modelName); ok {
		if typedRepo, ok := repo.(*GormRepository[M]); ok {
			return typedRepo, true
		}
	}
	return nil, false
}
