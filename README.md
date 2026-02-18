# Delazeza Forest

VALORANTストアチェッカー

## ビルド方法

### 1. フロントエンドをビルド

```sh
# フロントエンドのディレクトリに移動
cd frontend

# 依存関係インストール
bun install

# ビルド実行
bun run build

# プロジェクトルートに戻る
cd ..
```

### 2. 実行する

どちらかのコマンドを実行

```sh
# プロジェクトのルートディレクトリで実行する必要がある
go run ./backend/cmd/main.go

# 開発環境用実行コマンド
air
```

## 仕組み

### 1. 認証用URLを開く

```url
https://auth.riotgames.com/authorize?redirect_uri=https%3A%2F%2Fplayvalorant.com%2Fopt_in&client_id=play-valorant-web-prod&response_type=token%20id_token&nonce=1&scope=account%20openid
```

404と表示されるがリダイレクトされており、URLを確認するとアクセストークンなどが取得できている

### 2. URLを貼り付ける

リダイレクト後のURLからアクセストークンを読み取りユーザー情報にアクセスする

### 3. アクセスして取得したユーザーのストア情報を表示する
