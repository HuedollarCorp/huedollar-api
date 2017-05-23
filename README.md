# huedollar-api
HueDollar API, under the hood call the `http://api.promasters.net.br/cotacao/v1/valores` api

**API** 
https://huedollar-api.herokuapp.com/rates

## Response
```
{
  status: true,
  valores: {
    USD: {
      nome: "Dólar",
      valor: 3.2662,
      ultima_consulta: 1495569901,
      fonte: "UOL Economia - http://economia.uol.com.br/"
    }
  }
}
```

## Examples

### curl
```
curl https://huedollar-api.herokuapp.com/rates
```

### Browser
```
fetch('https://huedollar-api.herokuapp.com/rates')
  .then(r => r.json())
  .then(...)
```
