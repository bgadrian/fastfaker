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
| **&#123;hexcolor&#125;** | #v08d8x |
| **&#123;hipsterparagraphavg&#125;** | Beard typewriter goth taxidermy pinterest umami gluten-free narwhal freegan disrupt plaid bitters bicycle rights brooklyn keffiyeh meh bicycle rights flannel. Wayfarers brooklyn franzen hoodie skateboard mumblecore viral flannel etsy pork belly cardigan pinterest asymmetrical 8-bit DIY helvetica readymade pabst. Cold-pressed flexitarian tousled knausgaard fashion axe cray gastropub you probably haven't heard of them lomo heirloom vegan try-hard austin kickstarter gluten-free you probably haven't heard of them lo-fi viral. Cred tote bag echo mustache leggings kickstarter sustainable viral tilde semiotics cold-pressed selfies health celiac polaroid franzen chicharrones keffiyeh. 90's occupy carry chicharrones lomo twee chicharrones fixie blog neutra iPhone schlitz tacos brunch semiotics before they sold out whatever listicle. Paleo kickstarter williamsburg poutine meh cornhole scenester plaid seitan fixie migas dreamcatcher kombucha tofu biodiesel blue bottle pork belly whatever. VHS before they sold out fashion axe shabby chic vegan leggings distillery gentrify loko chillwave swag keffiyeh salvia photo booth hella cliche jean shorts portland. Paleo forage food truck narwhal vice salvia five dollar toast cornhole beard put a bird on it VHS gluten-free polaroid five dollar toast paleo cornhole hoodie retro. |
| **&#123;hipstersentenceavg&#125;** | Sartorial trust fund flannel ennui tumblr pop-up next level banh mi migas kogi vinyl swag franzen kombucha paleo pickled tofu Wes Anderson. |
| **&#123;hipsterword&#125;** | XOXO |
| **&#123;hour&#125;** | 12 |
| **&#123;httpmethod&#125;** | HEAD |
| **&#123;int16&#125;** | 11736 |
| **&#123;int32&#125;** | 600870614 |
| **&#123;int64&#125;** | -2060699869677000463 |
| **&#123;int64positive&#125;** | 8343085160255639013 |
| **&#123;int8&#125;** | 6 |
| **&#123;ipv4address&#125;** | 89.187.57.100 |
| **&#123;ipv6address&#125;** | 2001:cafe:8794:e9d4:69af:b33a:e471:4055 |
| **&#123;jobdescriptor&#125;** | Direct |
| **&#123;joblevel&#125;** | Implementation |
| **&#123;jobtitle&#125;** | Supervisor |
| **&#123;lastname&#125;** | Hartmann |
| **&#123;latitude&#125;** | 85.03619042973045 |
| **&#123;letter&#125;** | h |
| **&#123;longitude&#125;** | 2.728196828969317 |
| **&#123;md5&#125;** | 509138621c92d7b390ef48625c87f5b9 |
| **&#123;mimetype&#125;** | application/macbinary |
| **&#123;minute&#125;** | 36 |
| **&#123;month&#125;** | January |
| **&#123;name&#125;** | Monte Bashirian |
| **&#123;nameprefix&#125;** | Mr. |
| **&#123;namesuffix&#125;** | V |
| **&#123;nanosecond&#125;** | 208956050 |
| **&#123;operauseragent&#125;** | Opera/8.75 (Windows NT 5.1; en-US) Presto/2.13.193 Version/13.00 |
| **&#123;paragraphavg&#125;** | Et explicabo sed nulla exercitationem eveniet et eos eos voluptate inventore aliquam ut non consequatur exercitationem ut optio. Quis ratione quia quos ullam fugiat omnis nulla ut sed non et assumenda est ut quidem minima ea. Ex dolores enim maiores error rerum eum sint mollitia at ratione et ut et dolorum error magni autem. Dolorem deserunt molestiae voluptates excepturi quod voluptas eligendi ipsam nihil a quo maiores autem exercitationem nihil illum ipsum. Autem dolorum est modi architecto qui repellendus perferendis rem natus placeat corporis est explicabo asperiores consequatur doloribus error. In sint sed nisi quis cumque omnis a impedit reiciendis tempore unde dolores distinctio et odit qui et. Cupiditate odio corrupti et maiores placeat ratione dolor necessitatibus voluptatem enim pariatur occaecati perferendis doloremque qui voluptatem qui. Reprehenderit et molestiae aperiam animi non et reiciendis non molestias officiis aut ab molestias incidunt ullam qui impedit. |
| **&#123;passwordfull&#125;** | m$=4*@SQ?W&v |
| **&#123;phone&#125;** | 7962016519 |
| **&#123;phoneformatted&#125;** | 789.501.3948 |
| **&#123;safariuseragent&#125;** | Mozilla/5.0 (Windows; U; Windows NT 6.1) AppleWebKit/531.2.8 (KHTML, like Gecko) Version/4.1 Safari/531.2.8 |
| **&#123;safecolor&#125;** | yellow |
| **&#123;second&#125;** | 2 |
| **&#123;sentenceavg&#125;** | Enim laborum asperiores consectetur non nemo recusandae et explicabo accusamus et dolor enim et aut et repellat qui. |
| **&#123;sha256&#125;** | ff5bdebbbfc2abd2744a01c0501004bee94b00888e0f57feacbeb4e14cb2c649 |
| **&#123;simplestatuscode&#125;** | 200 |
| **&#123;ssn&#125;** | 840343837 |
| **&#123;state&#125;** | Virginia |
| **&#123;stateabr&#125;** | UT |
| **&#123;statuscode&#125;** | 404 |
| **&#123;street&#125;** | 15553 East Stravenuetown |
| **&#123;streetname&#125;** | Loaf |
| **&#123;streetnumber&#125;** | 55619 |
| **&#123;streetprefix&#125;** | North |
| **&#123;streetsuffix&#125;** | side |
| **&#123;timezone&#125;** | Mauritius Standard Time |
| **&#123;timezoneabv&#125;** | MST |
| **&#123;timezonefull&#125;** | (UTC+12:00) Fiji |
| **&#123;timezoneoffset&#125;** | 8 |
| **&#123;transmissiongeartype&#125;** | Automatic |
| **&#123;uint16&#125;** | 32253 |
| **&#123;uint32&#125;** | 1697838639 |
| **&#123;uint64&#125;** | 5681518370095422064 |
| **&#123;uint8&#125;** | 23 |
| **&#123;url&#125;** | http://www.leadmindshare.io/next-generation/granular/end-to-end/aggregate |
| **&#123;useragent&#125;** | Mozilla/5.0 (Windows NT 4.0) AppleWebKit/5361 (KHTML, like Gecko) Chrome/39.0.894.0 Mobile Safari/5361 |
| **&#123;username&#125;** | Crona6487 |
| **&#123;uuid&#125;** | c88804f6-e54f-476a-af4a-7f2360ab7333 |
| **&#123;vehicletype&#125;** | Van |
| **&#123;weekday&#125;** | Saturday |
| **&#123;word&#125;** | fugiat |
| **&#123;year&#125;** | 1975 |
| **&#123;zip&#125;** | 21859 |



> This file is generated automatically at build.

