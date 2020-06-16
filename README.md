# 概要

goでAPIを作成する際のシンプルな設計のサンプル。  
グロースや保守性を考えると品質はあまり良くないが、  
「参考に作ったらできた。なぜ動いてるかは謎」って状態を極力避けるために、あえて部品化は抑えて作成。

## 使用パッケージ

パッケージは、公式以外は基本使用しない。  
初心者だとgo言語の仕様なのか、パッケージの機能なのか判断難しいからね。  
下記に採用したパッケージと理由を記述。

### github.com/gin-gonic/gin  

HTTPメソッドの判断とルーティング処理を公式のみで書くと、不必要に冗長になるので、似たようなパッケージでechoやgolliraもあるのだが、githubのスター数が一番多いのでginを採用

### github.com/gin-contrib/cors

CORS設定で必要。net/httpでやろうとすると、ginを使用する場合記述がめんどくさいので採用。  
CORS設定できないとローカルでテストができない。
[CORSについて](https://developer.mozilla.org/ja/docs/Web/HTTP/CORS)

### github.com/go-sql-driver/mysql  

公式のデータベースパッケージを、mysqlで使用するために拡張してくれる。
