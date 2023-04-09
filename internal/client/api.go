package client

import (
	"bpzh_vk_bot/internal/config"
	"crypto/tls"
	"net/http"
)

type Client struct {
	client *http.Client
	config *config.ApiConfig
}

func NewClient(cfg *config.ApiConfig) *Client {
	return &Client{
		config: cfg,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

//func (c *Client) GetUserGroups(wbuserId, employeeId int64) (group []domain.Group, err error) {
//
//	bodyBytes, err := json.Marshal(body)
//	if err != nil {
//		log.Error().Err(err).Msg("Ошибка сборки rbac")
//		return
//	}
//
//	req, _ := http.NewRequest(http.MethodPost, c.config.URL+"/api/v1/groups/", bytes.NewBuffer(bodyBytes))
//
//	req.Header.Add("Authorization", c.config.Token)
//
//	resp, err := c.client.Do(req)
//	if err != nil {
//		log.Error().Err(err).Msg(fmt.Sprintf("ошибка проверки токена: %s", err))
//		err = fmt.Errorf("ошибка проверки токена: %s", err)
//		return
//	}
//	defer func() { _ = resp.Body.Close() }()
//
//	rawBody, err := io.ReadAll(resp.Body)
//	if err != nil {
//		log.Error().Err(err).Msg(fmt.Sprintf("ошибка чтения тела ответа rbac: %s", err))
//		err = fmt.Errorf("ошибка чтения тела ответа rbac: %s", err)
//		return
//	}
//
//	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
//		log.Error().Err(err).Msg(fmt.Sprintf("ошибка проверки токена rbac. Статус %d: %s", resp.StatusCode, rawBody))
//		err = fmt.Errorf("ошибка проверки токена rbac. Статус %d: %s", resp.StatusCode, rawBody)
//		return
//	}
//
//	type rbacResp struct {
//		Data struct {
//			Items []domain.Group `json:"items"`
//		} `json:"data"`
//	}
//
//	var r rbacResp
//	err = json.Unmarshal(rawBody, &r)
//	if err != nil {
//		log.Error().Err(err).Msg(fmt.Sprintf("ошибка разбора тела ответа rbac: %s", err))
//		err = fmt.Errorf("ошибка разбора тела ответа rbac: %s", err)
//		return
//	}
//
//	group = r.Data.Items
//	return
//}
