# FastFaker Templates

This document contains all the available variables, with examples that can be used with the Faker.Template, Faker.TemplateCustom and Faker.Struct methods.

A variable is an alias to the public Faker method with the same name, for example:
* **&#123;name&#125;** will be replaced by the result from Faker.Name() method
* **&#123;hackernoun&#125;** will be replaced by the result from Faker.HackerNoun() method

> Note: the Faker.Template() function has 2 more special characters: '#' and '?', they will be replaced by a digit or an ASCII lowercase letter. If you do want this functionality use Faker.TemplateCustom().

For more info see [TEMPLATES](./TEMPLATES.md).

| Variable | Example |
|---------|--------|
| **&#123;avatarurl&#125;** | http://pipsum.com/80x80.jpg |
| **&#123;beeralcohol&#125;** | 4.7% |
| **&#123;beerblg&#125;** | 9.3°Blg |
| **&#123;beerhop&#125;** | Glacier |
| **&#123;beeribu&#125;** | 42 IBU |
| **&#123;beermalt&#125;** | Victory |
| **&#123;beername&#125;** | Trappistes Rochefort 10 |
| **&#123;beerstyle&#125;** | Strong Ale |
| **&#123;beeryeast&#125;** | 2035 - American Lager |
| **&#123;binary&#125;** | 1 |
| **&#123;booltext&#125;** | false |
| **&#123;browser&#125;** | firefox |
| **&#123;bs&#125;** | plug-and-play |
| **&#123;buzzword&#125;** | Intuitive |
| **&#123;carmaker&#125;** | Skoda |
| **&#123;carmodel&#125;** | Mark Lt |
| **&#123;chromeuseragent&#125;** | Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_9_7) AppleWebKit/5311 (KHTML, like Gecko) Chrome/38.0.807.0 Mobile Safari/5311 |
| **&#123;city&#125;** | Alexaburgh |
| **&#123;color&#125;** | SteelBlue |
| **&#123;company&#125;** | Bauch-Koepp |
| **&#123;companysuffix&#125;** | Group |
| **&#123;country&#125;** | Palau |
| **&#123;countryabr&#125;** | CA |
| **&#123;creditcardcvv&#125;** | 369 |
| **&#123;creditcardexp&#125;** | 12/23 |
| **&#123;creditcardnumber&#125;** | 6564563828173105 |
| **&#123;creditcardnumberluhn&#125;** | 2221005338120163 |
| **&#123;creditcardtype&#125;** | MasterCard |
| **&#123;currencylong&#125;** | Malaysia Ringgit |
| **&#123;currencyshort&#125;** | BWP |
| **&#123;datecurrentyearstr&#125;** | 2019-10-02 21:33:26.12022322 +0000 UTC |
| **&#123;datestr&#125;** | 1977-05-25 20:37:07.645099301 +0000 UTC |
| **&#123;day&#125;** | 26 |
| **&#123;digit&#125;** | 1 |
| **&#123;domainname&#125;** | dynamicdistributed.info |
| **&#123;domainsuffix&#125;** | net |
| **&#123;email&#125;** | dantevandervort@borer.biz |
| **&#123;extension&#125;** | jar |
| **&#123;firefoxuseragent&#125;** | Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5 rv:7.0) Gecko/1930-10-10 Firefox/37.0 |
| **&#123;firstname&#125;** | Micah |
| **&#123;float32&#125;** | 8.799216e+37 |
| **&#123;float64&#125;** | 7.024500304461049e+306 |
| **&#123;float64unary&#125;** | 0.23227607590065746 |
| **&#123;fueltype&#125;** | LPG |
| **&#123;gender&#125;** | female |
| **&#123;hackerabbreviation&#125;** | FTP |
| **&#123;hackeradjective&#125;** | 1080p |
| **&#123;hackeringverb&#125;** | backing up |
| **&#123;hackernoun&#125;** | panel |
| **&#123;hackerphrase&#125;** | I'Ll synthesize the neural HTTP transmitter, that should connect the USB program! |
| **&#123;hackerverb&#125;** | generate |
| **&#123;hexcolor&#125;** | #8d7cb0 |
| **&#123;hipsterparagraphavg&#125;** | Locavore polaroid swag waistcoat aesthetic pitchfork beard typewriter goth taxidermy pinterest umami gluten-free narwhal freegan disrupt plaid bitters. Bicycle rights brooklyn keffiyeh meh bicycle rights flannel wayfarers brooklyn franzen hoodie skateboard mumblecore viral flannel etsy pork belly cardigan pinterest. Asymmetrical 8-bit DIY helvetica readymade pabst cold-pressed flexitarian tousled knausgaard fashion axe cray gastropub you probably haven't heard of them lomo heirloom vegan try-hard. Austin kickstarter gluten-free you probably haven't heard of them lo-fi viral cred tote bag echo mustache leggings kickstarter sustainable viral tilde semiotics cold-pressed selfies. Health celiac polaroid franzen chicharrones keffiyeh 90's occupy carry chicharrones lomo twee chicharrones fixie blog neutra iPhone schlitz. Tacos brunch semiotics before they sold out whatever listicle paleo kickstarter williamsburg poutine meh cornhole scenester plaid seitan fixie migas dreamcatcher. Kombucha tofu biodiesel blue bottle pork belly whatever VHS before they sold out fashion axe shabby chic vegan leggings distillery gentrify loko chillwave swag keffiyeh. Salvia photo booth hella cliche jean shorts portland paleo forage food truck narwhal vice salvia five dollar toast cornhole beard put a bird on it VHS gluten-free. |
| **&#123;hipstersentenceavg&#125;** | Polaroid five dollar toast paleo cornhole hoodie retro sartorial trust fund flannel ennui tumblr pop-up next level banh mi migas kogi vinyl swag. |
| **&#123;hipsterword&#125;** | franzen |
| **&#123;hour&#125;** | 12 |
| **&#123;httpmethod&#125;** | HEAD |
| **&#123;int16&#125;** | 1940 |
| **&#123;int32&#125;** | 1201764533 |
| **&#123;int64&#125;** | -7031973140882749355 |
| **&#123;int64positive&#125;** | 2287462722726645208 |
| **&#123;int8&#125;** | -15 |
| **&#123;ipv4address&#125;** | 63.44.164.128 |
| **&#123;ipv6address&#125;** | 2001:cafe:8fa2:41ec:e411:a8e7:2ffb:63d6 |
| **&#123;jobdescriptor&#125;** | Global |
| **&#123;joblevel&#125;** | Accountability |
| **&#123;jobtitle&#125;** | Specialist |
| **&#123;language&#125;** | Pali |
| **&#123;languageabbreviation&#125;** | kg |
| **&#123;lastname&#125;** | Connelly |
| **&#123;latitude&#125;** | 70.4370417828174 |
| **&#123;letter&#125;** | y |
| **&#123;longitude&#125;** | -100.10945887879839 |
| **&#123;macaddress&#125;** | 0a:24:4c:6b:a5:54 |
| **&#123;md5&#125;** | 26481ce629f5a7348686c9f306d40168 |
| **&#123;mimetype&#125;** | audio/x-gsm |
| **&#123;minute&#125;** | 59 |
| **&#123;month&#125;** | June |
| **&#123;name&#125;** | Mertie Pagac |
| **&#123;nameprefix&#125;** | Mr. |
| **&#123;namesuffix&#125;** | Jr. |
| **&#123;nanosecond&#125;** | 749087077 |
| **&#123;operauseragent&#125;** | Opera/8.67 (Macintosh; U; PPC Mac OS X 10_7_2; en-US) Presto/2.9.253 Version/11.00 |
| **&#123;paragraphavg&#125;** | Non consequatur exercitationem ut optio quis ratione quia quos ullam fugiat omnis nulla ut sed non et assumenda. Est ut quidem minima ea ex dolores enim maiores error rerum eum sint mollitia at ratione et ut. Et dolorum error magni autem dolorem deserunt molestiae voluptates excepturi quod voluptas eligendi ipsam nihil a quo maiores. Autem exercitationem nihil illum ipsum autem dolorum est modi architecto qui repellendus perferendis rem natus placeat corporis est. Explicabo asperiores consequatur doloribus error in sint sed nisi quis cumque omnis a impedit reiciendis tempore unde dolores. Distinctio et odit qui et cupiditate odio corrupti et maiores placeat ratione dolor necessitatibus voluptatem enim pariatur occaecati. Perferendis doloremque qui voluptatem qui reprehenderit et molestiae aperiam animi non et reiciendis non molestias officiis aut ab. Molestias incidunt ullam qui impedit quidem corrupti culpa suscipit ut aut neque accusamus dolorem dolorum minima dolor et. |
| **&#123;passwordfull&#125;** | 3ag1IK3@W9Yu |
| **&#123;phone&#125;** | 5898917183 |
| **&#123;phoneformatted&#125;** | 1-736-568-4722 |
| **&#123;programminglanguage&#125;** | ACT-III |
| **&#123;question&#125;** | Cold-pressed fingerstache trust fund? |
| **&#123;quote&#125;** | "Pinterest pour-over direct trade crucifix slow-carb 3 wolf moon." - Brad Wunsch |
| **&#123;safariuseragent&#125;** | Mozilla/5.0 (Windows; U; Windows NT 6.0) AppleWebKit/535.38.3 (KHTML, like Gecko) Version/4.1 Safari/535.38.3 |
| **&#123;safecolor&#125;** | teal |
| **&#123;second&#125;** | 51 |
| **&#123;sentenceavg&#125;** | Assumenda dolor hic omnis dolores iusto non aliquid quo voluptatibus aliquid voluptatem dignissimos aut dolorem eum at rerum. |
| **&#123;sha256&#125;** | c0a93914554b42e19ad22c348d9f47d96e0772d2c81f5d1745a2d423d32c6b52 |
| **&#123;simplestatuscode&#125;** | 400 |
| **&#123;ssn&#125;** | 112043252 |
| **&#123;state&#125;** | Wyoming |
| **&#123;stateabr&#125;** | NV |
| **&#123;statuscode&#125;** | 403 |
| **&#123;street&#125;** | 38732 Port Knollmouth |
| **&#123;streetname&#125;** | Parkway |
| **&#123;streetnumber&#125;** | 8432 |
| **&#123;streetprefix&#125;** | North |
| **&#123;streetsuffix&#125;** | view |
| **&#123;timezone&#125;** | British Summer Time |
| **&#123;timezoneabv&#125;** | IST |
| **&#123;timezonefull&#125;** | (UTC+02:00) Athens, Bucharest |
| **&#123;timezoneoffset&#125;** | 9 |
| **&#123;transmissiongeartype&#125;** | Manual |
| **&#123;uint16&#125;** | 52337 |
| **&#123;uint32&#125;** | 1865567344 |
| **&#123;uint64&#125;** | 6162679829796920209 |
| **&#123;uint8&#125;** | 4 |
| **&#123;url&#125;** | https://www.internationale-enable.biz/deploy/cutting-edge/convergence/platforms |
| **&#123;useragent&#125;** | Mozilla/5.0 (X11; Linux i686) AppleWebKit/5340 (KHTML, like Gecko) Chrome/37.0.855.0 Mobile Safari/5340 |
| **&#123;username&#125;** | Prohaska8143 |
| **&#123;uuid&#125;** | d003cc02-1eb7-401c-969f-df9e54b057cd |
| **&#123;vehicletype&#125;** | Van |
| **&#123;weekday&#125;** | Wednesday |
| **&#123;word&#125;** | doloremque |
| **&#123;year&#125;** | 1906 |
| **&#123;zip&#125;** | 69612 |



> This file is generated automatically at build.

