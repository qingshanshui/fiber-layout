package database

import (
	"wat.ink/layout/fiber/internal/model"
	"wat.ink/layout/fiber/pkg/logger"
	"wat.ink/layout/fiber/pkg/utils"
)

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate() {
	logger.Info("Starting database migration...")

	err := DB.AutoMigrate(
		&model.User{},
		// 在这里添加其他需要迁移的模型
	)

	if err != nil {
		logger.Fatal("Database migration failed", "error", err)
	}

	logger.Info("Database migration completed successfully")
}

// Seed 填充初始数据
func Seed() {
	logger.Info("Starting database seeding...")

	// 检查是否已经有管理员用户
	var count int64
	DB.Model(&model.User{}).Count(&count)
	if count == 0 {
		// 生成盐值和密码哈希
		salt := utils.GenerateSalt(6)
		password := "admin123" // 默认密码
		hashedPassword := utils.GeneratePassword(password, salt)

		// 创建默认管理员用户
		admin := &model.User{
			Username: "admin",
			Email:    "admin@example.com",
			Password: hashedPassword,
			Salt:     salt,
		}

		if err := DB.Create(admin).Error; err != nil {
			logger.Error("Failed to create admin user", 
				"error", err,
				"username", admin.Username,
			)
		} else {
			logger.Info("Created default admin user",
				"username", admin.Username,
				"email", admin.Email,
				"password", password, // 记录默认密码到日志，方便开发环境使用
			)
		}
	} else {
		logger.Info("Admin user already exists, skipping seed")
	}

	logger.Info("Database seeding completed")
} 