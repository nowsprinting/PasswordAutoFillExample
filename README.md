# Password AutoFill Example for Safari/iOS apps

## 概要

以下の検証用。GAE/Goにデプロイして検証に使用する

### Safariから、ルートディレクトリのログインフォーム

1. Safariから https://YOUR_PROJECT_ID.appspot.com/ を開く
1. 任意のID/Passwordでログイン
1. 次にログインフォームを開いたときにAuto Fillされるか確認


### Safariから、サブディレクトリのログインフォーム

1. Safariから https://YOUR_PROJECT_ID.appspot.com/sub/ を開く
1. ルートディレクトリのID/PasswordがAuto Fillされるか確認
1. 任意のID/Passwordでログイン
1. 次にログインフォームを開いたときにAuto Fillされるか確認


### アプリから

1. アプリのAssociated Domainsに https://YOUR_PROJECT_ID.appspot.com/ を設定
1. ログイン画面のID/Password用UITextFiledに`textContentType`を設定
1. ログイン画面を開いたときにID/PasswordがAuto Fillされるか確認（アプリからは保存はされない）
1. アプリ内の`WKWebView`でログインフォームを開き、ID/PasswordがAuto Fillされるか確認（これは保存も有効なはず？）

参考: [[iOS 11] 新機能の Password AutoFill を実装する](https://dev.classmethod.jp/smartphone/iphone/ios-11-password-autofill/)



## ビルドとデプロイ

1. `static_files/apple-app-site-association.json`に、検証に使うiOSアプリのTeam ID + BundleIdentifierを設定
1. Google Cloud PlatformコンソールでApp Engineの新規プロジェクトを作成し、下記コマンドでデプロイ

    $ gcloud config set project YOUR_PROJECT_ID
    $ gcloud app deploy app.yaml --version 1


## iOSの設定

iOSの設定でパスワード自動入力を許可している必要がある

- iOS 11以下: 設定 > Safari -> 自動入力 -> "ユーザー名とパスワード"をon
- iOS 12: 設定 -> パスワードとアカウント -> "パスワードを自動入力"をon
