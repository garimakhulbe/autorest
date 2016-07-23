package stringgroup

// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
)

// OperationsClient is the test Infrastructure for AutoRest Swagger BAT
type OperationsClient struct {
    ManagementClient
}
// NewOperationsClient creates an instance of the OperationsClient client.
func NewOperationsClient() OperationsClient {
  return OperationsClient{New()}
}

// GetEmpty get empty string value value ''
func (client OperationsClient) GetEmpty() (result String, err error) {
    req, err := client.GetEmptyPreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetEmpty", nil , "Failure preparing request")
    }

    resp, err := client.GetEmptySender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetEmpty", resp, "Failure sending request")
    }

    result, err = client.GetEmptyResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetEmpty", resp, "Failure responding to request")
    }

    return
}

// GetEmptyPreparer prepares the GetEmpty request.
func (client OperationsClient) GetEmptyPreparer() (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/string/empty"))
    return preparer.Prepare(&http.Request{})
}

// GetEmptySender sends the GetEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) GetEmptySender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetEmptyResponder handles the response to the GetEmpty request. The method always
// closes the http.Response Body.
func (client OperationsClient) GetEmptyResponder(resp *http.Response) (result String, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// GetMbcs get mbcs string value
// '啊齄丂狛狜隣郎隣兀﨩ˊ▇█〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€ '
func (client OperationsClient) GetMbcs() (result String, err error) {
    req, err := client.GetMbcsPreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetMbcs", nil , "Failure preparing request")
    }

    resp, err := client.GetMbcsSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetMbcs", resp, "Failure sending request")
    }

    result, err = client.GetMbcsResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetMbcs", resp, "Failure responding to request")
    }

    return
}

// GetMbcsPreparer prepares the GetMbcs request.
func (client OperationsClient) GetMbcsPreparer() (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/string/mbcs"))
    return preparer.Prepare(&http.Request{})
}

// GetMbcsSender sends the GetMbcs request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) GetMbcsSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetMbcsResponder handles the response to the GetMbcs request. The method always
// closes the http.Response Body.
func (client OperationsClient) GetMbcsResponder(resp *http.Response) (result String, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// GetNotProvided get String value when no string value is sent in response
// payload
func (client OperationsClient) GetNotProvided() (result String, err error) {
    req, err := client.GetNotProvidedPreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetNotProvided", nil , "Failure preparing request")
    }

    resp, err := client.GetNotProvidedSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetNotProvided", resp, "Failure sending request")
    }

    result, err = client.GetNotProvidedResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetNotProvided", resp, "Failure responding to request")
    }

    return
}

// GetNotProvidedPreparer prepares the GetNotProvided request.
func (client OperationsClient) GetNotProvidedPreparer() (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/string/notProvided"))
    return preparer.Prepare(&http.Request{})
}

// GetNotProvidedSender sends the GetNotProvided request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) GetNotProvidedSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetNotProvidedResponder handles the response to the GetNotProvided request. The method always
// closes the http.Response Body.
func (client OperationsClient) GetNotProvidedResponder(resp *http.Response) (result String, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// GetNull get null string value value
func (client OperationsClient) GetNull() (result String, err error) {
    req, err := client.GetNullPreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetNull", nil , "Failure preparing request")
    }

    resp, err := client.GetNullSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetNull", resp, "Failure sending request")
    }

    result, err = client.GetNullResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetNull", resp, "Failure responding to request")
    }

    return
}

// GetNullPreparer prepares the GetNull request.
func (client OperationsClient) GetNullPreparer() (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/string/null"))
    return preparer.Prepare(&http.Request{})
}

// GetNullSender sends the GetNull request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) GetNullSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetNullResponder handles the response to the GetNull request. The method always
// closes the http.Response Body.
func (client OperationsClient) GetNullResponder(resp *http.Response) (result String, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// GetWhitespace get string value with leading and trailing whitespace
// '<tab><space><space>Now is the time for all good men to come to the aid of
// their country<tab><space><space>'
func (client OperationsClient) GetWhitespace() (result String, err error) {
    req, err := client.GetWhitespacePreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetWhitespace", nil , "Failure preparing request")
    }

    resp, err := client.GetWhitespaceSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetWhitespace", resp, "Failure sending request")
    }

    result, err = client.GetWhitespaceResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "GetWhitespace", resp, "Failure responding to request")
    }

    return
}

// GetWhitespacePreparer prepares the GetWhitespace request.
func (client OperationsClient) GetWhitespacePreparer() (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/string/whitespace"))
    return preparer.Prepare(&http.Request{})
}

// GetWhitespaceSender sends the GetWhitespace request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) GetWhitespaceSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetWhitespaceResponder handles the response to the GetWhitespace request. The method always
// closes the http.Response Body.
func (client OperationsClient) GetWhitespaceResponder(resp *http.Response) (result String, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// PutEmpty set string value empty ''
//
// stringBody is possible values include: ''
func (client OperationsClient) PutEmpty(stringBody string) (result autorest.Response, err error) {
    req, err := client.PutEmptyPreparer(stringBody)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "PutEmpty", nil , "Failure preparing request")
    }

    resp, err := client.PutEmptySender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "PutEmpty", resp, "Failure sending request")
    }

    result, err = client.PutEmptyResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "PutEmpty", resp, "Failure responding to request")
    }

    return
}

// PutEmptyPreparer prepares the PutEmpty request.
func (client OperationsClient) PutEmptyPreparer(stringBody string) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPut(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/string/empty"),
                        autorest.WithJSON(stringBody))
    return preparer.Prepare(&http.Request{})
}

// PutEmptySender sends the PutEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) PutEmptySender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// PutEmptyResponder handles the response to the PutEmpty request. The method always
// closes the http.Response Body.
func (client OperationsClient) PutEmptyResponder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByClosing())
    result.Response = resp
    return
}

// PutMbcs set string value mbcs
// '啊齄丂狛狜隣郎隣兀﨩ˊ▇█〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€ '
//
// stringBody is possible values include:
// '啊齄丂狛狜隣郎隣兀﨩ˊ▇█〞〡￤℡㈱‐ー﹡﹢﹫、〓ⅰⅹ⒈€㈠㈩ⅠⅫ！￣ぁんァヶΑ︴АЯаяāɡㄅㄩ─╋︵﹄︻︱︳︴ⅰⅹɑɡ〇〾⿻⺁䜣€ '
func (client OperationsClient) PutMbcs(stringBody string) (result autorest.Response, err error) {
    req, err := client.PutMbcsPreparer(stringBody)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "PutMbcs", nil , "Failure preparing request")
    }

    resp, err := client.PutMbcsSender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "PutMbcs", resp, "Failure sending request")
    }

    result, err = client.PutMbcsResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "PutMbcs", resp, "Failure responding to request")
    }

    return
}

// PutMbcsPreparer prepares the PutMbcs request.
func (client OperationsClient) PutMbcsPreparer(stringBody string) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPut(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/string/mbcs"),
                        autorest.WithJSON(stringBody))
    return preparer.Prepare(&http.Request{})
}

// PutMbcsSender sends the PutMbcs request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) PutMbcsSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// PutMbcsResponder handles the response to the PutMbcs request. The method always
// closes the http.Response Body.
func (client OperationsClient) PutMbcsResponder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByClosing())
    result.Response = resp
    return
}

// PutNull set string value null
//
// stringBody is possible values include: ''
func (client OperationsClient) PutNull(stringBody string) (result autorest.Response, err error) {
    req, err := client.PutNullPreparer(stringBody)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "PutNull", nil , "Failure preparing request")
    }

    resp, err := client.PutNullSender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "PutNull", resp, "Failure sending request")
    }

    result, err = client.PutNullResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "PutNull", resp, "Failure responding to request")
    }

    return
}

// PutNullPreparer prepares the PutNull request.
func (client OperationsClient) PutNullPreparer(stringBody string) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPut(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/string/null"))
    if len(string(stringBody)) > 0 {
        preparer = autorest.DecoratePreparer(preparer,
                            autorest.WithJSON(stringBody))
    }
    return preparer.Prepare(&http.Request{})
}

// PutNullSender sends the PutNull request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) PutNullSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// PutNullResponder handles the response to the PutNull request. The method always
// closes the http.Response Body.
func (client OperationsClient) PutNullResponder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByClosing())
    result.Response = resp
    return
}

// PutWhitespace set String value with leading and trailing whitespace
// '<tab><space><space>Now is the time for all good men to come to the aid of
// their country<tab><space><space>'
//
// stringBody is possible values include: '    Now is the time for all good
// men to come to the aid of their country    '
func (client OperationsClient) PutWhitespace(stringBody string) (result autorest.Response, err error) {
    req, err := client.PutWhitespacePreparer(stringBody)
    if err != nil {
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "PutWhitespace", nil , "Failure preparing request")
    }

    resp, err := client.PutWhitespaceSender(req)
    if err != nil {
        result.Response = resp
        return result, autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "PutWhitespace", resp, "Failure sending request")
    }

    result, err = client.PutWhitespaceResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "stringgroup.OperationsClient", "PutWhitespace", resp, "Failure responding to request")
    }

    return
}

// PutWhitespacePreparer prepares the PutWhitespace request.
func (client OperationsClient) PutWhitespacePreparer(stringBody string) (*http.Request, error) {
    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPut(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPath("/string/whitespace"),
                        autorest.WithJSON(stringBody))
    return preparer.Prepare(&http.Request{})
}

// PutWhitespaceSender sends the PutWhitespace request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) PutWhitespaceSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// PutWhitespaceResponder handles the response to the PutWhitespace request. The method always
// closes the http.Response Body.
func (client OperationsClient) PutWhitespaceResponder(resp *http.Response) (result autorest.Response, err error) { 
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByClosing())
    result.Response = resp
    return
}

