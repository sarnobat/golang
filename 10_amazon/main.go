package main

import (
    "fmt"

    paapi5 "github.com/spiegel-im-spiegel/pa-api"
    "github.com/spiegel-im-spiegel/pa-api/entity"
    "github.com/spiegel-im-spiegel/pa-api/query"
)

func main() {
    //Create client
    client := paapi5.New(
        paapi5.WithMarketplace(paapi5.LocaleJapan),
    ).CreateClient(
        "mytag-20",
        "AKIAIOSFODNN7EXAMPLE",
        "1234567890",
    )

    //Make query
    q := query.NewGetItems(
        client.Marketplace(),
        client.PartnerTag(),
        client.PartnerType(),
    ).ASINs([]string{"B08TQSZYHN"}).EnableImages().EnableItemInfo().EnableParentASIN()

    //Requet and response
    body, err := client.Request(q)
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }
    //io.Copy(os.Stdout, bytes.NewReader(body))

    //Decode JSON
    res, err := entity.DecodeResponse(body)
    if err != nil {
        fmt.Printf("%+v\n", err)
        return
    }
    fmt.Println(res.String())
}

