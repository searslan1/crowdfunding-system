package user

// import (
// 	"KFS_Backend/pkg/logger"

// 	"fmt"

// 	"gorm.io/gorm"
// )

// type Repository struct {
//     db *gorm.DB
// }

// func NewRepository(db *gorm.DB) *Repository {
//     return &Repository{db: db}
// }

// func (r *Repository) CreateTables() error {
//     // Check table existence
//     exists, err := r.tableExists("users")
//     if (err != nil) {
//         return fmt.Errorf("error checking users table existence: %v", err)
//     }
//     if !exists {
//         // Drop table if exists
//         if err := r.db.Exec("DROP TABLE IF EXISTS users CASCADE").Error; err != nil {
//             return fmt.Errorf("failed to drop users table: %v", err)
//         }

//         // Create users table
//         err := r.db.Exec(`
//             CREATE TABLE users (
//                 user_id SERIAL PRIMARY KEY,
//                 email VARCHAR(255) UNIQUE NOT NULL,
//                 password_hash VARCHAR(255) NOT NULL,
//                 role VARCHAR(20) NOT NULL CHECK (role IN ('individual', 'corporate', 'admin', 'moderator')),
//                 verified BOOLEAN DEFAULT false,
//                 created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
//             )`).Error
//         if err != nil {
//             return fmt.Errorf("failed to create users table: %v", err)
//         }
//         logger.Info("✅ Users table created successfully")
//     }

//     exists, err = r.tableExists("auth_users")
//     if err != nil {
//         return fmt.Errorf("error checking auth_users table existence: %v", err)
//     }
//     if !exists {
//         // Drop table if exists
//         if err := r.db.Exec("DROP TABLE IF EXISTS auth_users CASCADE").Error; err != nil {
//             return fmt.Errorf("failed to drop auth_users table: %v", err)
//         }

//         // Create auth_users table
//         err := r.db.Exec(`
//             CREATE TABLE auth_users (
//                 user_id INTEGER PRIMARY KEY,
//                 failed_attempts INTEGER DEFAULT 0,
//                 account_locked_until TIMESTAMPTZ,
//                 created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
//                 updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
//                 FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
//             )`).Error
//         if err != nil {
//             return fmt.Errorf("failed to create auth_users table: %v", err)
//         }
//         logger.Info("✅ Auth users table created successfully")
//     }

//     exists, err = r.tableExists("user_sessions")
//     if err != nil {
//         return fmt.Errorf("error checking user_sessions table existence: %v", err)
//     }
//     if !exists {
//         // Create extension
//         if err := r.db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
//             return fmt.Errorf("failed to create uuid extension: %v", err)
//         }

//         // Drop table if exists
//         if err := r.db.Exec("DROP TABLE IF EXISTS user_sessions CASCADE").Error; err != nil {
//             return fmt.Errorf("failed to drop user_sessions table: %v", err)
//         }

//         // Create user_sessions table
//         err := r.db.Exec(`
//             CREATE TABLE user_sessions (
//                 session_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
//                 user_id INTEGER NOT NULL,
//                 ip_address INET NOT NULL,
//                 user_agent TEXT,
//                 device_info TEXT,
//                 login_time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
//                 logout_time TIMESTAMPTZ,
//                 is_active BOOLEAN DEFAULT true,
//                 last_activity TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
//                 refresh_token VARCHAR(255) UNIQUE,
//                 refresh_token_expiry TIMESTAMPTZ NOT NULL,
//                 FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
//             )`).Error
//         if err != nil {
//             return fmt.Errorf("failed to create user_sessions table: %v", err)
//         }
//         logger.Info("✅ User sessions table created successfully")
//     }

//     return nil
// }

// func (r *Repository) tableExists(tableName string) (bool, error) {
//     var exists bool
//     sqlDB, err := r.db.DB()
//     if err != nil {
//         return false, fmt.Errorf("failed to get underlying *sql.DB: %v", err)
//     }

//     err = sqlDB.QueryRow(`
//         SELECT EXISTS (
//             SELECT 1 FROM pg_tables
//             WHERE schemaname = 'public'
//             AND tablename = $1
//         )`, tableName).Scan(&exists)
//     return exists, err
// }
