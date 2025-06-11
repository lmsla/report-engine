Todo List

- [ ] 更新 `config/setting.yml`，加入 `auth` 區段
- [ ] 新增 `internal/infra/auth/provider.go`
- [ ] 重構現有 Keycloak code 到 `keycloak_provider.go`
- [ ] 建立 `radius_provider.go`，實作 RADIUS 認證
- [ ] 實作 `initAuthProvider()` 並注入 DI
- [ ] 更新 Middleware，改用通用介面
- [ ] 調整路由 `/login`、`/userinfo`
- [ ] 加入日誌與錯誤處理邏輯
- [ ] 撰寫單元測試 Stub Keycloak & Radius
- [ ] 更新 README，補充設定與使用說明
