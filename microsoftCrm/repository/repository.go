package repository

import "gorm.io/gorm"

type CrmRepositoryDB struct {
	CrmCollection *gorm.DB
}

func NewAccountsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}

func NewActivitiesRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewCasesRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewContactsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewKnowledgeArticlesRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewLeadsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewListsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewOpportunitiesRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewPingRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewProductsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewTasksRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewUsersRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
