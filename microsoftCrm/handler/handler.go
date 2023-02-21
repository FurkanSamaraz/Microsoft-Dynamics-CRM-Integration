package crm_handlers

import (
	"fmt"
	configs "main/config"
	handler_accounts "main/handler/accounts"
	handler_activities "main/handler/activities"
	crm_integration "main/package"
	repository "main/repository"
	crm_repository_accounts "main/repository/Accounts"
	crm_repository_activities "main/repository/Activities"

	handler_Cases "main/handler/Cases"
	crm_repository_Cases "main/repository/Cases"

	handler_Contacts "main/handler/Contacts"
	crm_repository_Contacts "main/repository/Contacts"

	handler_KnowledgeArticles "main/handler/Knowledge_Articles"
	crm_repository_KnowledgeArticles "main/repository/Knowledge_Articles"

	handler_Leads "main/handler/Leads"
	crm_repository_Leads "main/repository/Leads"

	handler_lists "main/handler/Lists"
	crm_repository_lists "main/repository/Lists"

	handler_Ping "main/handler/Ping"
	crm_repository_Ping "main/repository/Ping"

	handler_Products "main/handler/Products"
	crm_repository_Products "main/repository/Products"

	handler_Tasks "main/handler/Tasks"
	crm_repository_Tasks "main/repository/Tasks"

	handler_Users "main/handler/Users"
	crm_repository_Users "main/repository/Users"

	handler_Opportunities "main/handler/Opportunities"
	crm_repository_Opportunities "main/repository/Opportunities"

	service "main/service"
)

type CRM_CONTROLLER struct {
	Base *crm_integration.CRM_BASE
}

func App() {
	dbClient := configs.GetCollection()

	AccountsRepositoryDb := repository.NewAccountsRepositoryDB(dbClient)
	accounts := handler_accounts.AccountsHandler{Service: service.NewAccountsService(crm_repository_accounts.AccountsDB{CrmCollection: AccountsRepositoryDb.CrmCollection})}

	ActivitiesRepositoryDb := repository.NewActivitiesRepositoryDB(dbClient)
	activities := handler_activities.ActivitiesHandler{Service: service.NewActivitiesService(crm_repository_activities.ActivitiesDB{CrmCollection: ActivitiesRepositoryDb.CrmCollection})}

	CasesRepositoryDb := repository.NewCasesRepositoryDB(dbClient)
	Cases := handler_Cases.CasesHandler{Service: service.NewCasesService(crm_repository_Cases.CasesDB{CrmCollection: CasesRepositoryDb.CrmCollection})}

	ContactsRepositoryDb := repository.NewContactsRepositoryDB(dbClient)
	Contacts := handler_Contacts.ContactsHandler{Service: service.NewContactsService(crm_repository_Contacts.ContactsDB{CrmCollection: ContactsRepositoryDb.CrmCollection})}

	KnowledgeArticlesRepositoryDb := repository.NewKnowledgeArticlesRepositoryDB(dbClient)
	KnowledgeArticles := handler_KnowledgeArticles.Knowledge_ArticlesHandler{Service: service.NewKnowledgeArticlesService(crm_repository_KnowledgeArticles.KnowledgerArticlesDB{CrmCollection: KnowledgeArticlesRepositoryDb.CrmCollection})}

	LeadsRepositoryDb := repository.NewLeadsRepositoryDB(dbClient)
	Leads := handler_Leads.LeadsHandler{Service: service.NewLeadsService(crm_repository_Leads.LeadsDB{CrmCollection: LeadsRepositoryDb.CrmCollection})}

	listsRepositorysDb := repository.NewListsRepositoryDB(dbClient)
	Lists := handler_lists.ListsHandler{Service: service.NewListsService(crm_repository_lists.ListsDB{CrmCollection: listsRepositorysDb.CrmCollection})}

	PingRepositoryDb := repository.NewPingRepositoryDB(dbClient)
	Ping := handler_Ping.PingHandler{Service: service.NewPingService(crm_repository_Ping.PingDB{CrmCollection: PingRepositoryDb.CrmCollection})}

	ProductsRepositoryDb := repository.NewProductsRepositoryDB(dbClient)
	Products := handler_Products.ProductsHandler{Service: service.NewProductsService(crm_repository_Products.ProductsDB{CrmCollection: ProductsRepositoryDb.CrmCollection})}

	UsersRepositoryDb := repository.NewUsersRepositoryDB(dbClient)
	Users := handler_Users.UsersHandler{Service: service.NewUsersService(crm_repository_Users.UsersDB{CrmCollection: UsersRepositoryDb.CrmCollection})}

	TasksRepositoryDb := repository.NewTasksRepositoryDB(dbClient)
	Tasks := handler_Tasks.TasksHandler{Service: service.NewTasksService(crm_repository_Tasks.TasksDB{CrmCollection: TasksRepositoryDb.CrmCollection})}

	OpportunitiesRepositoryDb := repository.NewOpportunitiesRepositoryDB(dbClient)
	Opportunities := handler_Opportunities.OpportunitiesHandler{Service: service.NewOpportunitiesService(crm_repository_Opportunities.OpportunitiesDB{CrmCollection: OpportunitiesRepositoryDb.CrmCollection})}

	fmt.Println(
		accounts,
		activities,
		Cases,
		Contacts,
		KnowledgeArticles,
		Leads,
		Lists,
		Ping,
		Products,
		Users,
		Tasks,
		Opportunities,
	)

}
