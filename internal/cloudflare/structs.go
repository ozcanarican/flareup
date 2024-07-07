package cloudflare

import "time"

type RecordStruct struct {
	ID        string `json:"id"`
	ZoneID    string `json:"zone_id"`
	ZoneName  string `json:"zone_name"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	Proxiable bool   `json:"proxiable"`
	Proxied   bool   `json:"proxied"`
	TTL       int    `json:"ttl"`
	Locked    bool   `json:"locked"`
	Meta      struct {
		AutoAdded           bool `json:"auto_added"`
		ManagedByApps       bool `json:"managed_by_apps"`
		ManagedByArgoTunnel bool `json:"managed_by_argo_tunnel"`
	} `json:"meta,omitempty"`
	Comment    any       `json:"comment"`
	Tags       []any     `json:"tags"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
	Priority   int       `json:"priority,omitempty"`
	Meta0      struct {
		AutoAdded           bool `json:"auto_added"`
		EmailRouting        bool `json:"email_routing"`
		ManagedByApps       bool `json:"managed_by_apps"`
		ManagedByArgoTunnel bool `json:"managed_by_argo_tunnel"`
		ReadOnly            bool `json:"read_only"`
	} `json:"meta,omitempty"`
	Meta1 struct {
		AutoAdded           bool `json:"auto_added"`
		EmailRouting        bool `json:"email_routing"`
		ManagedByApps       bool `json:"managed_by_apps"`
		ManagedByArgoTunnel bool `json:"managed_by_argo_tunnel"`
		ReadOnly            bool `json:"read_only"`
	} `json:"meta,omitempty"`
	Meta2 struct {
		AutoAdded           bool `json:"auto_added"`
		EmailRouting        bool `json:"email_routing"`
		ManagedByApps       bool `json:"managed_by_apps"`
		ManagedByArgoTunnel bool `json:"managed_by_argo_tunnel"`
		ReadOnly            bool `json:"read_only"`
	} `json:"meta,omitempty"`
}

type RecordResponse struct {
	Result     []RecordStruct
	Success    bool  `json:"success"`
	Errors     []any `json:"errors"`
	Messages   []any `json:"messages"`
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		Count      int `json:"count"`
		TotalCount int `json:"total_count"`
		TotalPages int `json:"total_pages"`
	} `json:"result_info"`
}

type ZoneStruct struct {
	ID                  string    `json:"id"`
	Name                string    `json:"name"`
	Status              string    `json:"status"`
	Paused              bool      `json:"paused"`
	Type                string    `json:"type"`
	DevelopmentMode     int       `json:"development_mode"`
	NameServers         []string  `json:"name_servers"`
	OriginalNameServers []string  `json:"original_name_servers"`
	OriginalRegistrar   string    `json:"original_registrar"`
	OriginalDnshost     any       `json:"original_dnshost"`
	ModifiedOn          time.Time `json:"modified_on"`
	CreatedOn           time.Time `json:"created_on"`
	ActivatedOn         time.Time `json:"activated_on"`
	Meta                struct {
		Step                   int  `json:"step"`
		CustomCertificateQuota int  `json:"custom_certificate_quota"`
		PageRuleQuota          int  `json:"page_rule_quota"`
		PhishingDetected       bool `json:"phishing_detected"`
	} `json:"meta"`
	Owner struct {
		ID    any    `json:"id"`
		Type  string `json:"type"`
		Email any    `json:"email"`
	} `json:"owner"`
	Account struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"account"`
	Tenant struct {
		ID   any `json:"id"`
		Name any `json:"name"`
	} `json:"tenant"`
	TenantUnit struct {
		ID any `json:"id"`
	} `json:"tenant_unit"`
	Permissions []string `json:"permissions"`
	Plan        struct {
		ID                string `json:"id"`
		Name              string `json:"name"`
		Price             int    `json:"price"`
		Currency          string `json:"currency"`
		Frequency         string `json:"frequency"`
		IsSubscribed      bool   `json:"is_subscribed"`
		CanSubscribe      bool   `json:"can_subscribe"`
		LegacyID          string `json:"legacy_id"`
		LegacyDiscount    bool   `json:"legacy_discount"`
		ExternallyManaged bool   `json:"externally_managed"`
	} `json:"plan"`
}

type ZoneListResult struct {
	Result     []ZoneStruct
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		TotalPages int `json:"total_pages"`
		Count      int `json:"count"`
		TotalCount int `json:"total_count"`
	} `json:"result_info"`
	Success  bool  `json:"success"`
	Errors   []any `json:"errors"`
	Messages []any `json:"messages"`
}
