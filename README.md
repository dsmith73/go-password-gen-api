# go-password-gen-api

This endpoint will generate a random password in GO and return a json response, which can can be of varying length and include lowercase letters, uppercase letters, numbers, and symbols. Additionally, you can choose if the first character of the password should be lowercase, uppercase, either case, a number, or any (including symbols).  

<br>


## HTTP Default Request  

```sh
curl -X GET http://<base url>/password/generate
```

### Sample Response  
```json
{
    "password": "x+z84WHb:9L"
}
```  


## HTTP Declarative GET Request  

`/generate/{password length}/{lowercase letters}/{uppercase letters}/{numbers}/{symbols}/{first character}`  

```sh 
curl -X GET http://<base url>/password/generate/17/true/true/true/false/3
```

<br>

## HTTP Declarative POST Request  

`/generate`  

```sh 
curl -X POST -H "Content-Type: application/json" \
	-d '{"length": 23, "lowercase": true, "uppercase": true, "number": true, "symbol": true, "firstLetter": 5}' \
	http://<base url>/generate
```

<br>

---  

## Parameters  

| Parameter        | Options             | Default  | Description                                       |
| ---              | :---                | :---:    | ---                                               |
| Length           | 6 - 100             | **13**   | Amount of characters the password should contain  |
| Lowercase        | true / false        | true     | Include lowercase letters                         |
| Uppercase        | true / false        | true     | Include uppercase letters                         |
| Numbers          | true / false        | true     | Include numbers                                   |
| Symbols          | true / false        | true     | Include symbols                                   |
| First Character  | **1**.Lowercase<br> **2**.Uppercase<br> **3**.Both <br> **4**.Number <br> **5**.Any | **3** | Indicate if the password should begin with a lowercase, uppercase, or either, if it should begin with a number, or if it should begin with anything, including a symbol |



<br>

---

<br>

## Example - *Numbers only*  

```sh 
curl -X GET http://<base url>/password/generate/9/false/false/true/false/4
```


### Sample Numbers only Response  

```json
{
    "password": "388146707"
}
```  

