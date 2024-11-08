# TODO

## MUST

- controller 改修 

  - service に切り出す。

- repository 改修(済)

  - model を切り出す

    - thread と post、user と session は集約して、2 つのエンティティ(モデル)にまとめる。
    - `/model/user.go`と`/model/thread.go`に記述

  - repository をエンティティごとに分ける。

    - `repository/user.go`と`repository/thread.go`に記述
    - repository とのデータのやり取りはエンティティ(集約)単位で行う。

  - /service を作成し、model を用いてアプリのロジックを表現する。

  - sql クエリは全部 repository に封じ込める。

    - service では全てのデータをエンティティとして扱い、repository の save メソッドや find メソッドで永続化する。
    - repository には save, find, delete 系のメソッドを持たせる。
    - create メソッドはエンティティ(モデル)のメソッドとして実装する。create で生成されるのは構造体で、repository の save メソッドで永続化する。
    - また、update 系のメソッドはモデルのメソッドとして実装する。変更したドメインオブジェクトを repository の save メソッドで上書き保存する。

  - ディレクトリ構成: `/internal`内は、`/repository`,`/model`,`/service`,`/controller`の順に依存したパッケージの構成になる。
  - `/model`は何にも依存しない。`/repository`と`/service`の間のデータのやり取りに使う。

- repository のテスト

- CI 組む

## WANT

- logger 実装
- エラー処理まとめる
- スタイル当てる
