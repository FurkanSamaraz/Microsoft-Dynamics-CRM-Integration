package crm_structures

type CreateLeads struct {
	Attributes struct {
		Address1Addressid               string `json:"address1_addressid"`
		Address1Addresstypecode         int    `json:"address1_addresstypecode"`
		Address1City                    string `json:"address1_city"`
		Address1Composite               string `json:"address1_composite"`
		Address1Country                 string `json:"address1_country"`
		Address1County                  string `json:"address1_county"`
		Address1Fax                     string `json:"address1_fax"`
		Address1Latitude                int    `json:"address1_latitude"`
		Address1Line1                   string `json:"address1_line1"`
		Address1Line2                   string `json:"address1_line2"`
		Address1Line3                   string `json:"address1_line3"`
		Address1Longitude               int    `json:"address1_longitude"`
		Address1Name                    string `json:"address1_name"`
		Address1Postalcode              string `json:"address1_postalcode"`
		Address1Postofficebox           string `json:"address1_postofficebox"`
		Address1Shippingmethodcode      int    `json:"address1_shippingmethodcode"`
		Address1Stateorprovince         string `json:"address1_stateorprovince"`
		Address1Telephone1              string `json:"address1_telephone1"`
		Address1Telephone2              string `json:"address1_telephone2"`
		Address1Telephone3              string `json:"address1_telephone3"`
		Address1Upszone                 string `json:"address1_upszone"`
		Address1Utcoffset               int    `json:"address1_utcoffset"`
		Address2Addressid               string `json:"address2_addressid"`
		Address2Addresstypecode         int    `json:"address2_addresstypecode"`
		Address2City                    string `json:"address2_city"`
		Address2Composite               string `json:"address2_composite"`
		Address2Country                 string `json:"address2_country"`
		Address2County                  string `json:"address2_county"`
		Address2Fax                     string `json:"address2_fax"`
		Address2Latitude                int    `json:"address2_latitude"`
		Address2Line1                   string `json:"address2_line1"`
		Address2Line2                   string `json:"address2_line2"`
		Address2Line3                   string `json:"address2_line3"`
		Address2Longitude               int    `json:"address2_longitude"`
		Address2Name                    string `json:"address2_name"`
		Address2Postalcode              string `json:"address2_postalcode"`
		Address2Postofficebox           string `json:"address2_postofficebox"`
		Address2Shippingmethodcode      int    `json:"address2_shippingmethodcode"`
		Address2Stateorprovince         string `json:"address2_stateorprovince"`
		Address2Telephone1              string `json:"address2_telephone1"`
		Address2Telephone2              string `json:"address2_telephone2"`
		Address2Telephone3              string `json:"address2_telephone3"`
		Address2Upszone                 string `json:"address2_upszone"`
		Address2Utcoffset               int    `json:"address2_utcoffset"`
		Budgetamount                    int    `json:"budgetamount"`
		BudgetamountBase                int    `json:"budgetamount_base"`
		Budgetstatus                    int    `json:"budgetstatus"`
		Campaignid                      string `json:"campaignid"`
		Companyname                     string `json:"companyname"`
		Confirminterest                 bool   `json:"confirminterest"`
		Createdby                       string `json:"createdby"`
		Createdon                       int    `json:"createdon"`
		Createdonbehalfby               string `json:"createdonbehalfby"`
		CustomeridAccount               string `json:"customerid_account"`
		CustomeridContact               string `json:"customerid_contact"`
		Decisionmaker                   bool   `json:"decisionmaker"`
		Description                     string `json:"description"`
		Donotbulkemail                  bool   `json:"donotbulkemail"`
		Donotemail                      bool   `json:"donotemail"`
		Donotfax                        bool   `json:"donotfax"`
		Donotphone                      bool   `json:"donotphone"`
		Donotpostalmail                 bool   `json:"donotpostalmail"`
		Donotsendmm                     bool   `json:"donotsendmm"`
		Emailaddress1                   string `json:"emailaddress1"`
		Emailaddress2                   string `json:"emailaddress2"`
		Emailaddress3                   string `json:"emailaddress3"`
		Entityimage                     string `json:"entityimage"`
		EntityimageTimestamp            int    `json:"entityimage_timestamp"`
		EntityimageURL                  string `json:"entityimage_url"`
		Entityimageid                   string `json:"entityimageid"`
		Estimatedamount                 int    `json:"estimatedamount"`
		EstimatedamountBase             int    `json:"estimatedamount_base"`
		Estimatedclosedate              int    `json:"estimatedclosedate"`
		Estimatedvalue                  int    `json:"estimatedvalue"`
		Evaluatefit                     bool   `json:"evaluatefit"`
		Exchangerate                    int    `json:"exchangerate"`
		Fax                             string `json:"fax"`
		Firstname                       string `json:"firstname"`
		Followemail                     bool   `json:"followemail"`
		Fullname                        string `json:"fullname"`
		Importsequencenumber            int    `json:"importsequencenumber"`
		Industrycode                    int    `json:"industrycode"`
		Initialcommunication            int    `json:"initialcommunication"`
		Jobtitle                        string `json:"jobtitle"`
		Lastname                        string `json:"lastname"`
		Lastonholdtime                  int    `json:"lastonholdtime"`
		Lastusedincampaign              int    `json:"lastusedincampaign"`
		Leadid                          string `json:"leadid"`
		Leadqualitycode                 int    `json:"leadqualitycode"`
		Leadsourcecode                  int    `json:"leadsourcecode"`
		Masterid                        string `json:"masterid"`
		Merged                          bool   `json:"merged"`
		Middlename                      string `json:"middlename"`
		Mobilephone                     string `json:"mobilephone"`
		Modifiedby                      string `json:"modifiedby"`
		Modifiedon                      int    `json:"modifiedon"`
		Modifiedonbehalfby              string `json:"modifiedonbehalfby"`
		Need                            int    `json:"need"`
		Numberofemployees               int    `json:"numberofemployees"`
		Onholdtime                      int    `json:"onholdtime"`
		Originatingcaseid               string `json:"originatingcaseid"`
		Overriddencreatedon             int    `json:"overriddencreatedon"`
		Ownerid                         string `json:"ownerid"`
		Owningbusinessunit              string `json:"owningbusinessunit"`
		Owningteam                      string `json:"owningteam"`
		Owninguser                      string `json:"owninguser"`
		Pager                           string `json:"pager"`
		Parentaccountid                 string `json:"parentaccountid"`
		Parentcontactid                 string `json:"parentcontactid"`
		Participatesinworkflow          bool   `json:"participatesinworkflow"`
		Preferredcontactmethodcode      int    `json:"preferredcontactmethodcode"`
		Prioritycode                    int    `json:"prioritycode"`
		Processid                       string `json:"processid"`
		Purchaseprocess                 int    `json:"purchaseprocess"`
		Purchasetimeframe               int    `json:"purchasetimeframe"`
		Qualificationcomments           string `json:"qualificationcomments"`
		Qualifyingopportunityid         string `json:"qualifyingopportunityid"`
		Relatedobjectid                 string `json:"relatedobjectid"`
		Revenue                         int    `json:"revenue"`
		RevenueBase                     int    `json:"revenue_base"`
		Salesstage                      int    `json:"salesstage"`
		Salesstagecode                  int    `json:"salesstagecode"`
		Salutation                      string `json:"salutation"`
		SchedulefollowupProspect        int    `json:"schedulefollowup_prospect"`
		SchedulefollowupQualify         int    `json:"schedulefollowup_qualify"`
		Sic                             string `json:"sic"`
		SLALeadSLA                      string `json:"sla_lead_sla"`
		SlainvokedidLeadSLA             string `json:"slainvokedid_lead_sla"`
		Stageid                         string `json:"stageid"`
		StageidProcessstage             string `json:"stageid_processstage"`
		Statecode                       int    `json:"statecode"`
		Statuscode                      int    `json:"statuscode"`
		Subject                         string `json:"subject"`
		Telephone1                      string `json:"telephone1"`
		Telephone2                      string `json:"telephone2"`
		Telephone3                      string `json:"telephone3"`
		Timespentbymeonemailandmeetings string `json:"timespentbymeonemailandmeetings"`
		Timezoneruleversionnumber       int    `json:"timezoneruleversionnumber"`
		Transactioncurrencyid           string `json:"transactioncurrencyid"`
		Traversedpath                   string `json:"traversedpath"`
		Utcconversiontimezonecode       int    `json:"utcconversiontimezonecode"`
		Versionnumber                   int    `json:"versionnumber"`
		Websiteurl                      string `json:"websiteurl"`
		Yomicompanyname                 string `json:"yomicompanyname"`
		Yomifirstname                   string `json:"yomifirstname"`
		Yomifullname                    string `json:"yomifullname"`
		Yomilastname                    string `json:"yomilastname"`
		Yomimiddlename                  string `json:"yomimiddlename"`
	} `json:"attributes"`
	FetchMetaInfo bool `json:"fetchMetaInfo"`
}
type GetLeads struct {
	Attributes struct {
		Address1Addressid               string `json:"address1_addressid"`
		Address1Addresstypecode         int    `json:"address1_addresstypecode"`
		Address1City                    string `json:"address1_city"`
		Address1Composite               string `json:"address1_composite"`
		Address1Country                 string `json:"address1_country"`
		Address1County                  string `json:"address1_county"`
		Address1Fax                     string `json:"address1_fax"`
		Address1Latitude                int    `json:"address1_latitude"`
		Address1Line1                   string `json:"address1_line1"`
		Address1Line2                   string `json:"address1_line2"`
		Address1Line3                   string `json:"address1_line3"`
		Address1Longitude               int    `json:"address1_longitude"`
		Address1Name                    string `json:"address1_name"`
		Address1Postalcode              string `json:"address1_postalcode"`
		Address1Postofficebox           string `json:"address1_postofficebox"`
		Address1Shippingmethodcode      int    `json:"address1_shippingmethodcode"`
		Address1Stateorprovince         string `json:"address1_stateorprovince"`
		Address1Telephone1              string `json:"address1_telephone1"`
		Address1Telephone2              string `json:"address1_telephone2"`
		Address1Telephone3              string `json:"address1_telephone3"`
		Address1Upszone                 string `json:"address1_upszone"`
		Address1Utcoffset               int    `json:"address1_utcoffset"`
		Address2Addressid               string `json:"address2_addressid"`
		Address2Addresstypecode         int    `json:"address2_addresstypecode"`
		Address2City                    string `json:"address2_city"`
		Address2Composite               string `json:"address2_composite"`
		Address2Country                 string `json:"address2_country"`
		Address2County                  string `json:"address2_county"`
		Address2Fax                     string `json:"address2_fax"`
		Address2Latitude                int    `json:"address2_latitude"`
		Address2Line1                   string `json:"address2_line1"`
		Address2Line2                   string `json:"address2_line2"`
		Address2Line3                   string `json:"address2_line3"`
		Address2Longitude               int    `json:"address2_longitude"`
		Address2Name                    string `json:"address2_name"`
		Address2Postalcode              string `json:"address2_postalcode"`
		Address2Postofficebox           string `json:"address2_postofficebox"`
		Address2Shippingmethodcode      int    `json:"address2_shippingmethodcode"`
		Address2Stateorprovince         string `json:"address2_stateorprovince"`
		Address2Telephone1              string `json:"address2_telephone1"`
		Address2Telephone2              string `json:"address2_telephone2"`
		Address2Telephone3              string `json:"address2_telephone3"`
		Address2Upszone                 string `json:"address2_upszone"`
		Address2Utcoffset               int    `json:"address2_utcoffset"`
		Budgetamount                    int    `json:"budgetamount"`
		BudgetamountBase                int    `json:"budgetamount_base"`
		Budgetstatus                    int    `json:"budgetstatus"`
		Campaignid                      string `json:"campaignid"`
		Companyname                     string `json:"companyname"`
		Confirminterest                 bool   `json:"confirminterest"`
		Createdby                       string `json:"createdby"`
		Createdon                       int    `json:"createdon"`
		Createdonbehalfby               string `json:"createdonbehalfby"`
		CustomeridAccount               string `json:"customerid_account"`
		CustomeridContact               string `json:"customerid_contact"`
		Decisionmaker                   bool   `json:"decisionmaker"`
		Description                     string `json:"description"`
		Donotbulkemail                  bool   `json:"donotbulkemail"`
		Donotemail                      bool   `json:"donotemail"`
		Donotfax                        bool   `json:"donotfax"`
		Donotphone                      bool   `json:"donotphone"`
		Donotpostalmail                 bool   `json:"donotpostalmail"`
		Donotsendmm                     bool   `json:"donotsendmm"`
		Emailaddress1                   string `json:"emailaddress1"`
		Emailaddress2                   string `json:"emailaddress2"`
		Emailaddress3                   string `json:"emailaddress3"`
		Entityimage                     string `json:"entityimage"`
		EntityimageTimestamp            int    `json:"entityimage_timestamp"`
		EntityimageURL                  string `json:"entityimage_url"`
		Entityimageid                   string `json:"entityimageid"`
		Estimatedamount                 int    `json:"estimatedamount"`
		EstimatedamountBase             int    `json:"estimatedamount_base"`
		Estimatedclosedate              int    `json:"estimatedclosedate"`
		Estimatedvalue                  int    `json:"estimatedvalue"`
		Evaluatefit                     bool   `json:"evaluatefit"`
		Exchangerate                    int    `json:"exchangerate"`
		Fax                             string `json:"fax"`
		Firstname                       string `json:"firstname"`
		Followemail                     bool   `json:"followemail"`
		Fullname                        string `json:"fullname"`
		Importsequencenumber            int    `json:"importsequencenumber"`
		Industrycode                    int    `json:"industrycode"`
		Initialcommunication            int    `json:"initialcommunication"`
		Jobtitle                        string `json:"jobtitle"`
		Lastname                        string `json:"lastname"`
		Lastonholdtime                  int    `json:"lastonholdtime"`
		Lastusedincampaign              int    `json:"lastusedincampaign"`
		Leadid                          string `json:"leadid"`
		Leadqualitycode                 int    `json:"leadqualitycode"`
		Leadsourcecode                  int    `json:"leadsourcecode"`
		Masterid                        string `json:"masterid"`
		Merged                          bool   `json:"merged"`
		Middlename                      string `json:"middlename"`
		Mobilephone                     string `json:"mobilephone"`
		Modifiedby                      string `json:"modifiedby"`
		Modifiedon                      int    `json:"modifiedon"`
		Modifiedonbehalfby              string `json:"modifiedonbehalfby"`
		Need                            int    `json:"need"`
		Numberofemployees               int    `json:"numberofemployees"`
		Onholdtime                      int    `json:"onholdtime"`
		Originatingcaseid               string `json:"originatingcaseid"`
		Overriddencreatedon             int    `json:"overriddencreatedon"`
		Ownerid                         string `json:"ownerid"`
		Owningbusinessunit              string `json:"owningbusinessunit"`
		Owningteam                      string `json:"owningteam"`
		Owninguser                      string `json:"owninguser"`
		Pager                           string `json:"pager"`
		Parentaccountid                 string `json:"parentaccountid"`
		Parentcontactid                 string `json:"parentcontactid"`
		Participatesinworkflow          bool   `json:"participatesinworkflow"`
		Preferredcontactmethodcode      int    `json:"preferredcontactmethodcode"`
		Prioritycode                    int    `json:"prioritycode"`
		Processid                       string `json:"processid"`
		Purchaseprocess                 int    `json:"purchaseprocess"`
		Purchasetimeframe               int    `json:"purchasetimeframe"`
		Qualificationcomments           string `json:"qualificationcomments"`
		Qualifyingopportunityid         string `json:"qualifyingopportunityid"`
		Relatedobjectid                 string `json:"relatedobjectid"`
		Revenue                         int    `json:"revenue"`
		RevenueBase                     int    `json:"revenue_base"`
		Salesstage                      int    `json:"salesstage"`
		Salesstagecode                  int    `json:"salesstagecode"`
		Salutation                      string `json:"salutation"`
		SchedulefollowupProspect        int    `json:"schedulefollowup_prospect"`
		SchedulefollowupQualify         int    `json:"schedulefollowup_qualify"`
		Sic                             string `json:"sic"`
		SLALeadSLA                      string `json:"sla_lead_sla"`
		SlainvokedidLeadSLA             string `json:"slainvokedid_lead_sla"`
		Stageid                         string `json:"stageid"`
		StageidProcessstage             string `json:"stageid_processstage"`
		Statecode                       int    `json:"statecode"`
		Statuscode                      int    `json:"statuscode"`
		Subject                         string `json:"subject"`
		Telephone1                      string `json:"telephone1"`
		Telephone2                      string `json:"telephone2"`
		Telephone3                      string `json:"telephone3"`
		Timespentbymeonemailandmeetings string `json:"timespentbymeonemailandmeetings"`
		Timezoneruleversionnumber       int    `json:"timezoneruleversionnumber"`
		Transactioncurrencyid           string `json:"transactioncurrencyid"`
		Traversedpath                   string `json:"traversedpath"`
		Utcconversiontimezonecode       int    `json:"utcconversiontimezonecode"`
		Versionnumber                   int    `json:"versionnumber"`
		Websiteurl                      string `json:"websiteurl"`
		Yomicompanyname                 string `json:"yomicompanyname"`
		Yomifirstname                   string `json:"yomifirstname"`
		Yomifullname                    string `json:"yomifullname"`
		Yomilastname                    string `json:"yomilastname"`
		Yomimiddlename                  string `json:"yomimiddlename"`
	} `json:"attributes"`
	FetchMetaInfo bool   `json:"fetchMetaInfo"`
	ID            string `json:"id"`
}
type UpdateLeads struct {
	Attributes struct {
		Address1Addressid               string `json:"address1_addressid"`
		Address1Addresstypecode         int    `json:"address1_addresstypecode"`
		Address1City                    string `json:"address1_city"`
		Address1Composite               string `json:"address1_composite"`
		Address1Country                 string `json:"address1_country"`
		Address1County                  string `json:"address1_county"`
		Address1Fax                     string `json:"address1_fax"`
		Address1Latitude                int    `json:"address1_latitude"`
		Address1Line1                   string `json:"address1_line1"`
		Address1Line2                   string `json:"address1_line2"`
		Address1Line3                   string `json:"address1_line3"`
		Address1Longitude               int    `json:"address1_longitude"`
		Address1Name                    string `json:"address1_name"`
		Address1Postalcode              string `json:"address1_postalcode"`
		Address1Postofficebox           string `json:"address1_postofficebox"`
		Address1Shippingmethodcode      int    `json:"address1_shippingmethodcode"`
		Address1Stateorprovince         string `json:"address1_stateorprovince"`
		Address1Telephone1              string `json:"address1_telephone1"`
		Address1Telephone2              string `json:"address1_telephone2"`
		Address1Telephone3              string `json:"address1_telephone3"`
		Address1Upszone                 string `json:"address1_upszone"`
		Address1Utcoffset               int    `json:"address1_utcoffset"`
		Address2Addressid               string `json:"address2_addressid"`
		Address2Addresstypecode         int    `json:"address2_addresstypecode"`
		Address2City                    string `json:"address2_city"`
		Address2Composite               string `json:"address2_composite"`
		Address2Country                 string `json:"address2_country"`
		Address2County                  string `json:"address2_county"`
		Address2Fax                     string `json:"address2_fax"`
		Address2Latitude                int    `json:"address2_latitude"`
		Address2Line1                   string `json:"address2_line1"`
		Address2Line2                   string `json:"address2_line2"`
		Address2Line3                   string `json:"address2_line3"`
		Address2Longitude               int    `json:"address2_longitude"`
		Address2Name                    string `json:"address2_name"`
		Address2Postalcode              string `json:"address2_postalcode"`
		Address2Postofficebox           string `json:"address2_postofficebox"`
		Address2Shippingmethodcode      int    `json:"address2_shippingmethodcode"`
		Address2Stateorprovince         string `json:"address2_stateorprovince"`
		Address2Telephone1              string `json:"address2_telephone1"`
		Address2Telephone2              string `json:"address2_telephone2"`
		Address2Telephone3              string `json:"address2_telephone3"`
		Address2Upszone                 string `json:"address2_upszone"`
		Address2Utcoffset               int    `json:"address2_utcoffset"`
		Budgetamount                    int    `json:"budgetamount"`
		BudgetamountBase                int    `json:"budgetamount_base"`
		Budgetstatus                    int    `json:"budgetstatus"`
		Campaignid                      string `json:"campaignid"`
		Companyname                     string `json:"companyname"`
		Confirminterest                 bool   `json:"confirminterest"`
		Createdby                       string `json:"createdby"`
		Createdon                       int    `json:"createdon"`
		Createdonbehalfby               string `json:"createdonbehalfby"`
		CustomeridAccount               string `json:"customerid_account"`
		CustomeridContact               string `json:"customerid_contact"`
		Decisionmaker                   bool   `json:"decisionmaker"`
		Description                     string `json:"description"`
		Donotbulkemail                  bool   `json:"donotbulkemail"`
		Donotemail                      bool   `json:"donotemail"`
		Donotfax                        bool   `json:"donotfax"`
		Donotphone                      bool   `json:"donotphone"`
		Donotpostalmail                 bool   `json:"donotpostalmail"`
		Donotsendmm                     bool   `json:"donotsendmm"`
		Emailaddress1                   string `json:"emailaddress1"`
		Emailaddress2                   string `json:"emailaddress2"`
		Emailaddress3                   string `json:"emailaddress3"`
		Entityimage                     string `json:"entityimage"`
		EntityimageTimestamp            int    `json:"entityimage_timestamp"`
		EntityimageURL                  string `json:"entityimage_url"`
		Entityimageid                   string `json:"entityimageid"`
		Estimatedamount                 int    `json:"estimatedamount"`
		EstimatedamountBase             int    `json:"estimatedamount_base"`
		Estimatedclosedate              int    `json:"estimatedclosedate"`
		Estimatedvalue                  int    `json:"estimatedvalue"`
		Evaluatefit                     bool   `json:"evaluatefit"`
		Exchangerate                    int    `json:"exchangerate"`
		Fax                             string `json:"fax"`
		Firstname                       string `json:"firstname"`
		Followemail                     bool   `json:"followemail"`
		Fullname                        string `json:"fullname"`
		Importsequencenumber            int    `json:"importsequencenumber"`
		Industrycode                    int    `json:"industrycode"`
		Initialcommunication            int    `json:"initialcommunication"`
		Jobtitle                        string `json:"jobtitle"`
		Lastname                        string `json:"lastname"`
		Lastonholdtime                  int    `json:"lastonholdtime"`
		Lastusedincampaign              int    `json:"lastusedincampaign"`
		Leadid                          string `json:"leadid"`
		Leadqualitycode                 int    `json:"leadqualitycode"`
		Leadsourcecode                  int    `json:"leadsourcecode"`
		Masterid                        string `json:"masterid"`
		Merged                          bool   `json:"merged"`
		Middlename                      string `json:"middlename"`
		Mobilephone                     string `json:"mobilephone"`
		Modifiedby                      string `json:"modifiedby"`
		Modifiedon                      int    `json:"modifiedon"`
		Modifiedonbehalfby              string `json:"modifiedonbehalfby"`
		Need                            int    `json:"need"`
		Numberofemployees               int    `json:"numberofemployees"`
		Onholdtime                      int    `json:"onholdtime"`
		Originatingcaseid               string `json:"originatingcaseid"`
		Overriddencreatedon             int    `json:"overriddencreatedon"`
		Ownerid                         string `json:"ownerid"`
		Owningbusinessunit              string `json:"owningbusinessunit"`
		Owningteam                      string `json:"owningteam"`
		Owninguser                      string `json:"owninguser"`
		Pager                           string `json:"pager"`
		Parentaccountid                 string `json:"parentaccountid"`
		Parentcontactid                 string `json:"parentcontactid"`
		Participatesinworkflow          bool   `json:"participatesinworkflow"`
		Preferredcontactmethodcode      int    `json:"preferredcontactmethodcode"`
		Prioritycode                    int    `json:"prioritycode"`
		Processid                       string `json:"processid"`
		Purchaseprocess                 int    `json:"purchaseprocess"`
		Purchasetimeframe               int    `json:"purchasetimeframe"`
		Qualificationcomments           string `json:"qualificationcomments"`
		Qualifyingopportunityid         string `json:"qualifyingopportunityid"`
		Relatedobjectid                 string `json:"relatedobjectid"`
		Revenue                         int    `json:"revenue"`
		RevenueBase                     int    `json:"revenue_base"`
		Salesstage                      int    `json:"salesstage"`
		Salesstagecode                  int    `json:"salesstagecode"`
		Salutation                      string `json:"salutation"`
		SchedulefollowupProspect        int    `json:"schedulefollowup_prospect"`
		SchedulefollowupQualify         int    `json:"schedulefollowup_qualify"`
		Sic                             string `json:"sic"`
		SLALeadSLA                      string `json:"sla_lead_sla"`
		SlainvokedidLeadSLA             string `json:"slainvokedid_lead_sla"`
		Stageid                         string `json:"stageid"`
		StageidProcessstage             string `json:"stageid_processstage"`
		Statecode                       int    `json:"statecode"`
		Statuscode                      int    `json:"statuscode"`
		Subject                         string `json:"subject"`
		Telephone1                      string `json:"telephone1"`
		Telephone2                      string `json:"telephone2"`
		Telephone3                      string `json:"telephone3"`
		Timespentbymeonemailandmeetings string `json:"timespentbymeonemailandmeetings"`
		Timezoneruleversionnumber       int    `json:"timezoneruleversionnumber"`
		Transactioncurrencyid           string `json:"transactioncurrencyid"`
		Traversedpath                   string `json:"traversedpath"`
		Utcconversiontimezonecode       int    `json:"utcconversiontimezonecode"`
		Versionnumber                   int    `json:"versionnumber"`
		Websiteurl                      string `json:"websiteurl"`
		Yomicompanyname                 string `json:"yomicompanyname"`
		Yomifirstname                   string `json:"yomifirstname"`
		Yomifullname                    string `json:"yomifullname"`
		Yomilastname                    string `json:"yomilastname"`
		Yomimiddlename                  string `json:"yomimiddlename"`
	} `json:"attributes"`
	FetchMetaInfo bool `json:"fetchMetaInfo"`
}