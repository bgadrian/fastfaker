make bench OR go test -bench=. -benchmem 
goos: darwin  
goarch: amd64  
pkg: github.com/bgadrian/fastfaker/faker  
Table generated with tablesgenerator.com/markdown_tables 

| Benchmark              | Ops   | CPU  | MEM   | MEM alloc  |
|---------------------------------|-----------|-------------|------------|--------------|
| BenchmarkStreet-4                         	| 834050                                    	| 1432 ns/op  	| 62 B/op   	| 3 allocs/op  	|
| BenchmarkStreetNumber-4                   	| 2075772                                   	| 580 ns/op   	| 36 B/op   	| 2 allocs/op  	|
| BenchmarkStreetPrefix-4                   	| 5211439                                   	| 224 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkStreetName-4                     	| 5469650                                   	| 220 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkStreetSuffix-4                   	| 5169009                                   	| 224 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkCity-4                           	| 1876627                                   	| 639 ns/op   	| 15 B/op   	| 1 allocs/op  	|
| BenchmarkState-4                          	| 5134915                                   	| 220 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkStateAbr-4                       	| 5379502                                   	| 224 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkZip-4                            	| 2515968                                   	| 475 ns/op   	| 5 B/op    	| 1 allocs/op  	|
| BenchmarkCountry-4                        	| 4891246                                   	| 232 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkCountryAbr-4                     	| 5325213                                   	| 226 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkLatitude-4                       	| 72118604                                  	| 16.5 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkLongitude-4                      	| 72230757                                  	| 16.9 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkLatitudeInRange-4                	| 44111890                                  	| 27.5 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkLongitudeInRange-4               	| 43706727                                  	| 28.4 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkBeerName-4                       	| 5399041                                   	| 191 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkBeerStyle-4                      	| 4854436                                   	| 239 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkBeerHop-4                        	| 6081080                                   	| 198 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkBeerYeast-4                      	| 5776946                                   	| 205 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkBeerMalt-4                       	| 5307602                                   	| 239 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkBeerIbu-4                        	| 9327277                                   	| 122 ns/op   	| 8 B/op    	| 1 allocs/op  	|
| BenchmarkBeerAlcohol-4                    	| 1784971                                   	| 677 ns/op   	| 40 B/op   	| 3 allocs/op  	|
| BenchmarkBeerBlg-4                        	| 1760458                                   	| 681 ns/op   	| 48 B/op   	| 3 allocs/op  	|
| BenchmarkBool-4                           	| 49157251                                  	| 24.3 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkBoolText-4                       	| 49665682                                  	| 24.3 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkBinary-4                         	| 49269094                                  	| 24.5 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkColor-4                          	| 6126792                                   	| 197 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkSafeColor-4                      	| 7117698                                   	| 169 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkHexColor-4                       	| 4996138                                   	| 241 ns/op   	| 8 B/op    	| 1 allocs/op  	|
| BenchmarkRGBColor-4                       	| 9288019                                   	| 129 ns/op   	| 32 B/op   	| 1 allocs/op  	|
| BenchmarkCompany-4                        	| 1725397                                   	| 693 ns/op   	| 22 B/op   	| 1 allocs/op  	|
| BenchmarkCompanySuffix-4                  	| 7153242                                   	| 168 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkBuzzWord-4                       	| 6093070                                   	| 189 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkBS-4                             	| 6171804                                   	| 195 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkContact-4                        	| 706395                                    	| 1732 ns/op  	| 146 B/op  	| 6 allocs/op  	|
| BenchmarkPhone-4                          	| 8622144                                   	| 138 ns/op   	| 16 B/op   	| 1 allocs/op  	|
| BenchmarkPhoneFormatted-4                 	| 1812614                                   	| 663 ns/op   	| 16 B/op   	| 1 allocs/op  	|
| BenchmarkEmail-4                          	| 826363                                    	| 1458 ns/op  	| 98 B/op   	| 4 allocs/op  	|
| BenchmarkCurrency-4                       	| 5262478                                   	| 232 ns/op   	| 32 B/op   	| 1 allocs/op  	|
| BenchmarkCurrencyShort-4                  	| 6894070                                   	| 176 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkCurrencyLong-4                   	| 6230736                                   	| 184 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkPrice-4                          	| 58635384                                  	| 20.6 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkDate-4                           	| 2523836                                   	| 459 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkDateCurrentYear-4                	| 2947888                                   	| 413 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkDateRange-4                      	| 2843472                                   	| 421 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkMonth-4                          	| 27957531                                  	| 43.3 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkDay-4                            	| 31413968                                  	| 38.6 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkWeekDay-4                        	| 28319086                                  	| 43.3 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkYear-4                           	| 6522738                                   	| 186 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkHour-4                           	| 31536241                                  	| 38.3 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkMinute-4                         	| 29778072                                  	| 38.5 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkSecond-4                         	| 31509411                                  	| 38.0 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkNanoSecond-4                     	| 29342071                                  	| 41.0 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkTimeZone-4                       	| 6103005                                   	| 198 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkTimeZoneFull-4                   	| 5400486                                   	| 233 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkTimeZoneAbv-4                    	| 6024212                                   	| 191 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkTimeZoneOffset-4                 	| 4426400                                   	| 266 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkNewSafeFaker_Parallel-4          	| 609369                                    	| 2162 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkNewFastFaker_Parallel-4          	| 6928526                                   	| 177 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkMimeType-4                       	| 6927506                                   	| 172 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkExtension-4                      	| 5938406                                   	| 208 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkHackerPhrase-4                   	| 103494                                    	| 11863 ns/op 	| 1017 B/op 	| 13 allocs/op 	|
| BenchmarkHackerAbbreviation-4             	| 6563973                                   	| 186 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkHackerAdjective-4                	| 6329066                                   	| 190 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkHackerNoun-4                     	| 6237880                                   	| 193 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkHackerVerb-4                     	| 3716479                                   	| 319 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkHackerIngverb-4                  	| 4252633                                   	| 287 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkHipsterWord-4                    	| 6374347                                   	| 180 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkHipsterSentence-4                	| 585363                                    	| 2129 ns/op  	| 401 B/op  	| 5 allocs/op  	|
| BenchmarkHipsterParagraph-4               	| 46834                                     	| 25515 ns/op 	| 5285 B/op 	| 4 allocs/op  	|
| BenchmarkAvatarURL-4                      	| 3216836                                   	| 372 ns/op   	| 48 B/op   	| 3 allocs/op  	|
| BenchmarkImageURL-4                       	| 3182900                                   	| 378 ns/op   	| 48 B/op   	| 3 allocs/op  	|
| BenchmarkDomainName-4                     	| 1208848                                   	| 993 ns/op   	| 53 B/op   	| 2 allocs/op  	|
| BenchmarkDomainSuffix-4                   	| 6823706                                   	| 175 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkURL-4                            	| 561624                                    	| 2282 ns/op  	| 226 B/op  	| 4 allocs/op  	|
| BenchmarkHTTPMethod-4                     	| 6426768                                   	| 188 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkIPv4Address-4                    	| 1623061                                   	| 746 ns/op   	| 48 B/op   	| 5 allocs/op  	|
| BenchmarkIPv6Address-4                    	| 1213712                                   	| 987 ns/op   	| 96 B/op   	| 7 allocs/op  	|
| BenchmarkMacAddress-4                     	| 672650                                    	| 1806 ns/op  	| 79 B/op   	| 6 allocs/op  	|
| BenchmarkUsername-4                       	| 2420996                                   	| 543 ns/op   	| 16 B/op   	| 2 allocs/op  	|
| BenchmarkJob-4                            	| 868230                                    	| 1406 ns/op  	| 86 B/op   	| 2 allocs/op  	|
| BenchmarkJobTitle-4                       	| 6958218                                   	| 166 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkJobDescriptor-4                  	| 6499477                                   	| 178 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkJobLevel-4                       	| 5905996                                   	| 203 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkLanguage-4                       	| 4477191                                   	| 261 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkLanguageAbbreviation-4           	| 4457119                                   	| 265 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkProgrammingLanguage-4            	| 4454802                                   	| 270 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkLogLevel-4                       	| 6399550                                   	| 188 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkMap-4                            	| 257392                                    	| 4719 ns/op  	| 549 B/op  	| 6 allocs/op  	|
| BenchmarkReplaceWithNumbers-4             	| 3234552                                   	| 369 ns/op   	| 32 B/op   	| 1 allocs/op  	|
| BenchmarkNumerify-4                       	| 2679508                                   	| 461 ns/op   	| 16 B/op   	| 1 allocs/op  	|
| BenchmarkHexify-4                         	| 5169279                                   	| 244 ns/op   	| 8 B/op    	| 1 allocs/op  	|
| BenchmarkName-4                           	| 2326766                                   	| 520 ns/op   	| 17 B/op   	| 1 allocs/op  	|
| BenchmarkFirstName-4                      	| 6321150                                   	| 187 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkLastName-4                       	| 6208520                                   	| 193 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkNamePrefix-4                     	| 7242378                                   	| 166 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkNameSuffix-4                     	| 6084700                                   	| 198 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkNumber-4                         	| 38278070                                  	| 31.3 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkUint8-4                          	| 42217311                                  	| 28.5 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkUint16-4                         	| 41986030                                  	| 28.6 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkUint32-4                         	| 42758412                                  	| 28.0 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkUint64-4                         	| 25417172                                  	| 47.1 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkInt8-4                           	| 40714650                                  	| 28.6 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkInt16-4                          	| 40573660                                  	| 29.4 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkInt32-4                          	| 40262797                                  	| 28.3 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkInt64-4                          	| 24238872                                  	| 49.8 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkFloat32-4                        	| 48891394                                  	| 24.4 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkFloat32Range-4                   	| 48884260                                  	| 24.4 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkFloat64-4                        	| 58764861                                  	| 20.3 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkFloat64Range-4                   	| 60437208                                  	| 19.9 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkShuffleInts-4                    	| 4668235                                   	| 255 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkPassword-4                       	| 1000000                                   	| 1024 ns/op  	| 304 B/op  	| 6 allocs/op  	|
| BenchmarkCreditCard-4                     	| 786140                                    	| 1540 ns/op  	| 88 B/op   	| 4 allocs/op  	|
| BenchmarkCreditCardType-4                 	| 7094682                                   	| 169 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkCreditCardNumber-4               	| 1494565                                   	| 811 ns/op   	| 16 B/op   	| 1 allocs/op  	|
| BenchmarkCreditCardNumberLuhn-4           	| 141961                                    	| 8398 ns/op  	| 160 B/op  	| 10 allocs/op 	|
| BenchmarkCreditCardExp-4                  	| 5488742                                   	| 218 ns/op   	| 5 B/op    	| 1 allocs/op  	|
| BenchmarkCreditCardCvv-4                  	| 7334521                                   	| 165 ns/op   	| 3 B/op    	| 1 allocs/op  	|
| BenchmarkSSN-4                            	| 8190106                                   	| 147 ns/op   	| 16 B/op   	| 1 allocs/op  	|
| BenchmarkGender-4                         	| 45746403                                  	| 24.3 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkPerson-4                         	| 123134                                    	| 9837 ns/op  	| 736 B/op  	| 24 allocs/op 	|
| BenchmarkSimpleStatusCode-4               	| 9087405                                   	| 126 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkStatusCode-4                     	| 9112227                                   	| 132 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkLetter-4                         	| 31537593                                  	| 38.1 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkDigit-4                          	| 31696008                                  	| 38.1 ns/op  	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkLexify-4                         	| 4427620                                   	| 277 ns/op   	| 8 B/op    	| 1 allocs/op  	|
| BenchmarkShuffleStrings-4                 	| 5825574                                   	| 209 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkFaker_Template-4                 	| 228901                                    	| 4643 ns/op  	| 532 B/op  	| 11 allocs/op 	|
| BenchmarkFaker_TemplateCustom-4           	| 275032                                    	| 4415 ns/op  	| 467 B/op  	| 9 allocs/op  	|
| BenchmarkFaker_TemplateCustomConcurrent-4 	| 174292                                    	| 6900 ns/op  	| 1174 B/op 	| 24 allocs/op 	|
| BenchmarkUUID-4                           	| 6038590                                   	| 199 ns/op   	| 48 B/op   	| 1 allocs/op  	|
| BenchmarkUserAgent-4                      	| 567118                                    	| 2157 ns/op  	| 305 B/op  	| 5 allocs/op  	|
| BenchmarkChromeUserAgent-4                	| 756466                                    	| 1601 ns/op  	| 188 B/op  	| 5 allocs/op  	|
| BenchmarkFirefoxUserAgent-4               	| 420068                                    	| 2876 ns/op  	| 386 B/op  	| 7 allocs/op  	|
| BenchmarkSafariUserAgent-4                	| 513426                                    	| 2491 ns/op  	| 551 B/op  	| 7 allocs/op  	|
| BenchmarkOperaUserAgent-4                 	| 711361                                    	| 1772 ns/op  	| 216 B/op  	| 5 allocs/op  	|
| BenchmarkFaker_CarMaker-4                 	| 5195347                                   	| 227 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkFaker_CarModel-4                 	| 4716396                                   	| 253 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkFaker_Vehicle-4                  	| 853218                                    	| 1419 ns/op  	| 96 B/op   	| 1 allocs/op  	|
| BenchmarkFaker_VehicleType-4              	| 6693523                                   	| 181 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkWord-4                           	| 7101261                                   	| 171 ns/op   	| 0 B/op    	| 0 allocs/op  	|
| BenchmarkSentence-4                       	| 638743                                    	| 2023 ns/op  	| 325 B/op  	| 4 allocs/op  	|
| BenchmarkParagraph-4                      	| 45721                                     	| 26958 ns/op 	| 4795 B/op 	| 4 allocs/op  	|
| BenchmarkQuestion-4                       	| 643509                                    	| 1956 ns/op  	| 410 B/op  	| 6 allocs/op  	|
| BenchmarkQuote-4                          	| 544170                                    	| 2393 ns/op  	| 364 B/op  	| 5 allocs/op  	|
