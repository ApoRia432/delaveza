package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	REGION = "ap"
)

func GetStorefront(c echo.Context) error {
	puuid := c.QueryParam("puuid")
	// 外部リクエストを作成
	fmt.Println("https://pd." + REGION + ".a.pvp.net/store/v3/storefront/" + puuid)
	fmt.Println("query storefront by " + puuid)
	req, err := http.NewRequest("POST", "https://pd."+REGION+".a.pvp.net/store/v3/storefront/"+puuid, bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return c.String(http.StatusInternalServerError, "リクエストの作成に失敗しました")
	}

	req.Header.Set("X-Riot-ClientPlatform", "ew0KCSJwbGF0Zm9ybVR5cGUiOiAiUEMiLA0KCSJwbGF0Zm9ybU9TIjogIldpbmRvd3MiLA0KCSJwbGF0Zm9ybU9TVmVyc2lvbiI6ICIxMC4wLjE5MDQyLjEuMjU2LjY0Yml0IiwNCgkicGxhdGZvcm1DaGlwc2V0IjogIlVua25vd24iDQp9")
	req.Header.Set("X-Riot-ClientVersion", "release-12.02-shipping-9-4226954")
	// 元のリクエストのヘッダーをコピー（必要に応じて調整）
	req.Header.Set("X-Riot-Entitlements-JWT", c.Request().Header.Get("X-Riot-Entitlements-JWT"))
	req.Header.Set("Authorization", c.Request().Header.Get("Authorization"))

	// HTTPクライアントでリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.String(http.StatusBadGateway, "外部APIへの接続に失敗しました")
	}
	defer resp.Body.Close()

	// レスポンスボディを文字列で取得
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(bodyBytes)

	// 外部レスポンスをそのままクライアントへ返す
	fmt.Println("body: " + bodyString)
	c.Response().Header().Add("Content-Type", "application/json")
	return c.String(resp.StatusCode, bodyString)
}
