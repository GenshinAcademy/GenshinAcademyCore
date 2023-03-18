package models

type Url string

func CreateUrl(url string) (Url, error) {
    var err = VerifyUrl(url)
    if err != nil {
        return "", err
    }

    return Url(url), nil
}

func VerifyUrl(url string) error {
    //TODO
    return nil
}