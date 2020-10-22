package persistence

import (
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lupguo/go-ddd-sample/domain/entity"
	"github.com/lupguo/go-ddd-sample/domain/repository"
	"time"
)

// Repositories 总仓储机构提，包含多个领域仓储接口，以及一个DB实例
type Repositories struct {
	UploadImg repository.UploadImgRepo
	db        *gorm.DB
}

// NewRepositories 初始化所有域的总仓储实例，将实例通过依赖注入方式，将DB实例注入到领域层
func NewRepositories(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	cfg := &mysql.Config{
		User:                 DbUser,
		Passwd:               DbPassword,
		Net:                  "tcp",
		Addr:                 DbHost + ":" + DbPort,
		DBName:               DbName,
		Collation:            "utf8mb4_general_ci",
		Loc:                  time.FixedZone("Asia/Shanghai", 8*60*60),
		Timeout:              time.Second,
		ReadTimeout:          30 * time.Second,
		WriteTimeout:         30 * time.Second,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	// DBSource := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, "tcp", DbHost, DbPort, DbName)
	db, err := gorm.Open(DbDriver, cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	// 初始化总仓储实例
	return &Repositories{
		UploadImg: NewUploadImgPersis(db),
		db:        db,
	}, nil
}

// closes the database connection
func (s *Repositories) Close() error {
	return s.db.Close()
}

// This migrate all tables
func (s *Repositories) AutoMigrate() error {
	return s.db.AutoMigrate(&entity.UploadImg{}).Error
}
