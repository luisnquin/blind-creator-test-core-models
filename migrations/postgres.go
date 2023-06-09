package migrations

import (
	"context"
	"log"
	"time"

	"github.com/luisnquin/blind-creator-test-core-models/models/agencies"
	"github.com/luisnquin/blind-creator-test-core-models/models/campaign_creator_social_network_actions"
	"github.com/luisnquin/blind-creator-test-core-models/models/campaigns"
	"github.com/luisnquin/blind-creator-test-core-models/models/companies"
	"github.com/luisnquin/blind-creator-test-core-models/models/creator_social_networks"
	"github.com/luisnquin/blind-creator-test-core-models/models/creators"
	"github.com/luisnquin/blind-creator-test-core-models/models/user_agency_relations"
	"github.com/luisnquin/blind-creator-test-core-models/models/users"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SilentLogger struct{}

func (l SilentLogger) LogMode(logger.LogLevel) logger.Interface {
	return l
}

func (l SilentLogger) Info(context.Context, string, ...interface{}) {}

func (l SilentLogger) Warn(context.Context, string, ...interface{}) {}

func (l SilentLogger) Error(context.Context, string, ...interface{}) {}

func (l SilentLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
}

func ApplyMigrations(dbConnection *gorm.DB) {
	// if !applyMigrations {
	// 	return
	// }

	dbConnection.Logger = SilentLogger{} // Silent mode for now

	// MODELS
	log.Println("users...")
	dbConnection.AutoMigrate(&users.User{})

	log.Println("agencies...")
	dbConnection.AutoMigrate(&agencies.Agency{})

	log.Println("user_agency_relations...")
	dbConnection.AutoMigrate(&user_agency_relations.UserAgencyRelation{})

	log.Println("creators...")
	dbConnection.AutoMigrate(&creators.Creator{})

	log.Println("creator_social_networks...")
	dbConnection.AutoMigrate(&creator_social_networks.CreatorSocialNetworkAccount{})

	log.Println("companies...")
	dbConnection.AutoMigrate(&companies.Company{})

	log.Println("campaigns...")
	dbConnection.AutoMigrate(&campaigns.Campaign{})

	log.Println("campaign_creator_social_network_actions...")
	dbConnection.AutoMigrate(&campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions{})

	// RELATIONS

	// TODO: add foreign keys using model tags

	log.Println("user_agency_relations...")
	dbConnection.Model(&user_agency_relations.UserAgencyRelation{}) //.
	// AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").
	// AddForeignKey("agency_id", "agencies(id)", "CASCADE", "CASCADE")

	log.Println("creator_social_networks...")
	dbConnection.Model(&creator_social_networks.CreatorSocialNetworkAccount{}) //.
	// AddForeignKey("creator_id", "creators(id)", "CASCADE", "CASCADE")

	log.Println("companies...")
	dbConnection.Model(&companies.Company{}) //.
	// AddForeignKey("agency_id", "agencies(id)", "CASCADE", "CASCADE").
	// AddForeignKey("manager_id", "users(id)", "CASCADE", "CASCADE")

	log.Println("campaigns...")
	dbConnection.Model(&campaigns.Campaign{}) //.
	// AddForeignKey("agency_id", "agencies(id)", "CASCADE", "CASCADE").
	// AddForeignKey("manager_id", "users(id)", "CASCADE", "CASCADE").
	// AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")

	log.Println("campaign_creator_social_network_actions ...")
	dbConnection.Model(&campaign_creator_social_network_actions.CampaignCreatorSocialNetworkActions{}) //.
	// AddForeignKey("campaign_id", "campaigns(id)", "CASCADE", "CASCADE").
	// AddForeignKey("creator_id", "creators(id)", "CASCADE", "CASCADE").
	// AddForeignKey("creator_social_network_id", "creator_social_networks(id)", "CASCADE", "CASCADE")
}
