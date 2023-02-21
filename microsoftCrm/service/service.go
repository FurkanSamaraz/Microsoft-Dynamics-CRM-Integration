package crm_services

import (
	crm_repository_Accounts "main/repository/Accounts"
	crm_service_Accounts "main/service/Accounts"

	crm_repository_Activities "main/repository/Activities"
	crm_service_Activities "main/service/Activities"

	crm_repository_Cases "main/repository/Cases"
	crm_service_Cases "main/service/Cases"

	crm_repository_Contacts "main/repository/Contacts"
	crm_service_Contacts "main/service/Contacts"

	crm_repository_KnowledgeArticles "main/repository/Knowledge_Articles"
	crm_service_KnowledgeArticles "main/service/Knowledge_Articles"

	crm_repository_Leads "main/repository/Leads"
	crm_service_Leads "main/service/Leads"

	crm_repository_Lists "main/repository/Lists"
	crm_service_Lists "main/service/Lists"

	crm_repository_Opportunities "main/repository/Opportunities"
	crm_service_Opportunities "main/service/Opportunities"

	crm_repository_Ping "main/repository/Ping"
	crm_service_Ping "main/service/Ping"

	crm_repository_Products "main/repository/Products"
	crm_service_Products "main/service/Products"

	crm_repository_Tasks "main/repository/Tasks"
	crm_service_Tasks "main/service/Tasks"

	crm_repository_Users "main/repository/Users"
	crm_service_Users "main/service/Users"
)

func NewAccountsService(Repo crm_repository_Accounts.AccountsRepositoryInterface) crm_service_Accounts.AccountsCRMService {
	return crm_service_Accounts.AccountsCRMService{Repo: Repo}
}

func NewActivitiesService(Repo crm_repository_Activities.ActivitiesRepositoryInterface) crm_service_Activities.ActivitiesCRMService {
	return crm_service_Activities.ActivitiesCRMService{Repo: Repo}
}

func NewCasesService(Repo crm_repository_Cases.CasesRepositoryInterface) crm_service_Cases.CasesCRMService {
	return crm_service_Cases.CasesCRMService{Repo: Repo}
}

func NewContactsService(Repo crm_repository_Contacts.ContactsRepositoryInterface) crm_service_Contacts.ContactsCRMService {
	return crm_service_Contacts.ContactsCRMService{Repo: Repo}
}
func NewKnowledgeArticlesService(Repo crm_repository_KnowledgeArticles.KnowledgerArticlesRepositoryInterface) crm_service_KnowledgeArticles.KnowledgeArticlesCRMService {
	return crm_service_KnowledgeArticles.KnowledgeArticlesCRMService{Repo: Repo}
}
func NewLeadsService(Repo crm_repository_Leads.LeadsRepositoryInterface) crm_service_Leads.LeadsCRMService {
	return crm_service_Leads.LeadsCRMService{Repo: Repo}
}
func NewListsService(Repo crm_repository_Lists.ListsRepositoryInterface) crm_service_Lists.ListsCRMService {
	return crm_service_Lists.ListsCRMService{Repo: Repo}
}
func NewOpportunitiesService(Repo crm_repository_Opportunities.OpportunitiesRepositoryInterface) crm_service_Opportunities.OpportunitiesCRMService {
	return crm_service_Opportunities.OpportunitiesCRMService{Repo: Repo}
}
func NewProductsService(Repo crm_repository_Products.ProductsRepositoryInterface) crm_service_Products.ProductsCRMService {
	return crm_service_Products.ProductsCRMService{Repo: Repo}
}

func NewTasksService(Repo crm_repository_Tasks.TasksRepositoryInterface) crm_service_Tasks.TasksCRMService {
	return crm_service_Tasks.TasksCRMService{Repo: Repo}
}
func NewPingService(Repo crm_repository_Ping.PingRepositoryInterface) crm_service_Ping.PingCRMService {
	return crm_service_Ping.PingCRMService{Repo: Repo}
}
func NewUsersService(Repo crm_repository_Users.UsersRepositoryInterface) crm_service_Users.UsersCRMService {
	return crm_service_Users.UsersCRMService{Repo: Repo}
}
