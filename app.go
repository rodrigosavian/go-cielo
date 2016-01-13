package main



import (
    "fmt"
    "net/http"
    "net/url"
    "io/ioutil"
    "bytes"
    "encoding/xml"
)

const endpoint string = "https://qasecommerce.cielo.com.br/servicos/ecommwsec.do"
const appkey string = "1006993069"
const secret string = "25fbb99741c739dd84d7b06ec78c9bac718838630f30b112d033ce2e621b34f3"

type Transaction struct {
    XMLName xml.Name `xml:"requisicao-transacao"`
    Id string `xml:"id,attr"`
    Version string `xml:"versao,attr"`
    AppKey string `xml:"dados-ec>numero"`
    Secret string `xml:"dados-ec>chave"`
    Number string `xml:"dados-portador>numero"`
    DateUntil string `xml:"dados-portador>validade"`
    Indicator int `xml:"dados-portador>indicador"`
    SecurityId int `xml:"dados-portador>codigo-seguranca"`
    Token string `xml:"dados-portador>token"`
    OrderNumber string `xml:"dados-pedido>numero"`
    Value int `xml:"dados-pedido>valor"`
    Moeda int `xml:"dados-pedido>moeda"`
    CreatedAt string `xml:"dados-pedido>data-hora"`
    Description string `xml:"dados-pedido>descricao"`
    Language string `xml:"dados-pedido>idioma"`
    Bandeira string `xml:"forma-pagamento>bandeira"`
    Product string `xml:"forma-pagamento>produto"`
    Parcel int `xml:"forma-pagamento>parcelas"`
    ReturnUrl string `xml:"url-retorno"`
    Authorization int `xml:"autorizar"`
    Capture bool `xml:"capturar"`
}

func request(data string) (resp *http.Response, err error) {
    client := &http.Client{}
    body := url.Values{}
    body.Add("mensagem", data)

    req, err := http.NewRequest("POST", endpoint, bytes.NewBufferString(body.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    return client.Do(req)
}

func main() {
    t := Transaction{
        Id: "a97ab62a-7956-41ea-b03f-c2e9f612c293",
        Version: "1.2.1", 
        AppKey: appkey, 
        Secret: secret, 
        Number: "4012001037141112", 
        DateUntil: "052018", 
        Indicator: 1, 
        SecurityId: 123, 
        Token: "", 
        OrderNumber: "178148599", 
        Value: 1000, 
        Moeda: 986, 
        CreatedAt: "2011-12-07T11:43:37", 
        Description: "[origem:10.50.54.156]", 
        Language: "PT", 
        Bandeira: "visa", 
        Product: "A", 
        Parcel: 1, 
        ReturnUrl: "http://localhost/lojaexemplo/retorno.jsp", 
        Authorization: 1, 
        Capture: false}

    output, _ := xml.MarshalIndent(t, "  ", "    ")
    data := xml.Header + string(output)
    resp, _ := request(data)
    defer resp.Body.Close()
    content, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(content))
}
