## Project Setup
```
 git clone https://github.com/WarrenU/shake-search.git
 cd shake-search
 go run ./main.go
```
Server should start in a few seconds and say something like:
`2020/08/14 16:51:51 Starting server on Port 9000`
After it says that, you make make a request to the server. See CURL request example below:

## cURL Request Example
`curl http://localhost:9000/autocomplete?term=th`
API responses are json objects. The objects include an amount of words matched and words in an array. See Results for various terms below.

## Results For Various Query Params
 ### th
   ```
   {
      "amount":["25"],
      "words":[
         "the","that","this","thou","thy","thee","they","then","their","them","than","there","these","th","think","thus","though","therefore","those","thine", "thats","theres","three","thought","thing"
      ]
   }
   ```
 ### fr
   ```
   {
      "amount":["25"],
      "words":[
         "from","france","friends","friend","french","free","friar","fresh","freely","francis","frame","frown","friendship","fruit","friendly","frederick", "freedom","fright","froth","front","fran","frenchman","frowns","fray","frail"
      ]
   }
   ```
 ### pi
   ```
   {
      "amount":["25"],
      "words":[
         "pity","pistol","pisanio","piece","picture","pieces","pinch","pitch","pitiful","pierce","pit","piteous","pindarus","pine","pin","pitied","pickd","pillow","pilgrimage","pick","pisa","pipe","pigeons","piercing","pins"
      ]
   }
   ```
 ### sh
   ```
   {
      "amount":["25"],
      "words":[
         "shall","she","should","show","shame","shalt","shakespeare","shallow","shes","shepherd","shylock","shows","shake","short","shape","shadow","shouldst","sharp","shut","shore","showd","ship","shed","shortly","shell"
      ]
   }
   ```
 ### wu
   ```
   {
      "amount":["1"],
      "words":["wul"]
   }
   ```

 ### ar
   ```
   {
      "amount":["25"],
      "words":[
         "are","art","arms","arm","armado","army","ariel","archbishop","argument","arviragus","arthur","armour","armd","arise","armed","arrest","articles","arrivd","article","arrant","array","arch","arras","arts","arrows"
      ]
   }
   ```
 ### il
   ```
   {
      "amount":["25"],
      "words":[
         "ill","illinois","ild","illyria","il","ills","illfavourd","ilion","illustrious","ilium","illbeseeming","illfavouredly","illusion","illboding","illusions","illegitimate","illdisposd","illustrate","illtemperd","illwell","illstarrd","ilbow","illseeming","illbreeding","illumineth"
      ]
   }
   ```
 ### ne
   ```
   {
      "amount":["25"],
      "words":[
         "never","news","near","neer","new","neither","need","next","needs","neck","nerissa","nestor","neighbour","necessity","nephew","newly","needful","neighbours","nest","neglect","nearer","necessary","ned","negligence","nell"
      ]
   }
   ```
 ### se
   ```
   {
      "amount":["25"],
      "words":[
         "see","second","set","service","servant","sent","seen","send","seek","sea","serve","seem","sebastian","self","sense","servants","seems","senator","several","seven","seal","secret","senators","seat","seeming"
      ]
   }
   ```
 ### pl
   ```
   {
      "amount":["25"],
      "words":[
         "place","please","play","pleasure","plain","pluck","plague","plantagenet","pleasd","plot","plead","playd","places","plays","pleasures","players","pleasant","pluckd","plebeians","plant","pleasing","playing","plainly","pledge","pleases"
      ]
   }
   ```

## References
https://golangcode.com/how-to-remove-all-non-alphanumerical-characters-from-a-string/

https://github.com/golang/go/wiki/SliceTricks

https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
