# Numbers Everyone Should Know

## 単位

| 時間 | 比率 |
| --- | --- |
| s | $1$ |
| ms | $10^{-3}$ |
| μs | $10^{-6}$ |
| ns | $10^{-9}$ |
| ps | $10^{-12}$ |

| バイト | 2進数 | 10進数 |
| --- | --- | --- |
| PB | $2^{50}$ | $10^{15}$ |
| TB | $2^{40}$ | $10^{12}$ |
| GB | $2^{30}$ | $10^9$ |
| MB | $2^{20}$ | $10^6$ |
| KB  | $2^{10}$ | $10^3$ |
| B | $1$ | $1$ |

## Jeff Dean Numbers

| 操作 | ns | μs | ms | 補足 |
| --- | --- | --- | --- | --- |
| 通常のCPU命令 | 1 ns | | | |
| L1キャッシュ参照 | 0.5 ns | | | |
| 分岐予測ミス | 5 ns | | | |
| L2キャッシュ参照 | 7 ns | | | L1キャッシュの14倍 |
| Mutex lock/unlock | 25 ns | | | |
| メインメモリ参照 | 100 ns | | | L1キャッシュの200倍、L2キャッシュの20倍 |
| 1KBのデータを[Zippy](https://en.wikipedia.org/wiki/Snappy_(compression))で圧縮 | 3,000 ns | 3 μs | | |
| 1KBのデータを1Gbpsのネットワークで送信 | 10,000 ns | 10 μs | | |
| SSDから4KBをランダムにRead | 150,000 ns | 150 μs | | ~1GB/s SSD |
| メモリから1MBをシーケンシャルにRead | 250,000 ns | 250 μs | | |
| 同じデータセンター内のRTT | 500,000 ns | 500 μs | | |
| SSDから1MBをシーケンシャルにRead | 1,000,000 ns | 1,000 μs | 1ms | ~1GB/s SSD, メモリの4倍 |
| ディスクシーク | 10,000,000 ns | 10,000 μs | 10 ms | データセンター内のラウンドトリップの20倍 |
| ディスクから1MBをシーケンシャルにRead | 20,000,000 ns | 20,000 μs | 20 ms | メモリの80倍、SSDの20倍 |
| カルフォルニア-オランダ間のRTT | 150,000,000 ns | 150,000 μs | 150 ms | |

- 現在ではアメリカ西海岸-東京間のRTTは100msぐらい[^1]らしいので、今はカルフォルニア-オランダ間のRTTは80msぐらいでは？
- デフォルトゲートウェイへの通信速度は無線だと1, 2ms、有線だと1, 2μsぐらい

## 言語

| 言語 | 速度比率 | |
| --- | --- | --- |
| C/C++ | 1 | 1～10億ステップ/s |
| Go, Java, C# | 3 | |
| Python | 50~100 | |

| 型（C/C++） | バイト |
| --- | --- |
| bool | 1 B |
| char | 1 B |
| short | 2 B |
| int | 4 B |
| long | 4, 8 B |
| long long | 8 B |
| float | 4 B |
| double | 8 B |

- デフォルトのコールスタックの深さは、Linuxだと8MB、Pythonだと1000個、Javaだと１万個ぐらい

## コンピュータ

| | サイズ |
| --- | --- |
| L1キャッシュ | 32~128 KB |
| L2キャッシュ | 512 KB ~ 4 MB |
| L3キャッシュ | 16~256 MB |
| メインメモリ | 4~128 GB |
| 補助記憶装置 | 64 GB ~ 4 TB |

| | スループット |
| --- | --- |
| PCIe 4.0 | 5 GB/s |
| PCIe 5.0 | 10 GB/s |

| | 時間 |
| --- | --- |
| CPU周波数 | 1~3 GHz |
| CPUコア数 | 4~64 |
| 関数呼び出し | 1 ns |
| システムコール | 500 ns |
| コンテキストスイッチ | 10 μs |
| DB（MySQL, Redis）クエリ | 500 μs |

- HDDやSSDの平均寿命は約5年、年間故障率（AFR）は0.5~1.5%

## グラフィック

| | 値 |
| --- | --- |
| ピクセルサイズ | 3 B |
| 4K解像度 | 3840x2160 |
| HD解像度 | 1920x1080 |
| リフレッシュレート（フレームレート） | 通常：24, 30, 60 Hz(fps)、ゲーム：120, 144, 240 Hz(fps) |

## システムデザイン

- Webページの表示速度は数msが望ましく、2~3秒が限界
- １台のWebサーバーの最大同時接続数は200~1000ぐらい
- １台のDBサーバーで扱えるレコード数は数十万ぐらい、100万レコードを超えてくると厳しくなる

## 参照

[Latency Numbers Every Programmer Should Know](https://gist.github.com/jboner/2841832)
[Latency Numbers Every Programmer Should Know - comment](https://gist.github.com/jboner/2841832?permalink_comment_id=3707733#gistcomment-3707733)
["Numbers Everyone Should Know" from Jeff Dean](https://brenocon.com/dean_perf.html)
[Original "Numbers Everyone Should Know"](http://norvig.com/21-days.html#answers)
[Napkin Math](https://github.com/sirupsen/napkin-math)
[Speed comparision](https://github.com/niklas-heer/speed-comparison)
[The Computer Language Benchmarks Game](https://benchmarksgame-team.pages.debian.net/benchmarksgame/)

[^1]: [Yahoo! JAPAN アメリカデータセンタとネットワーク変遷 - 39ページ](https://www.janog.gr.jp/meeting/janog52/wp-content/uploads/2023/06/janog52-atpo-fukazawa.pdf)
