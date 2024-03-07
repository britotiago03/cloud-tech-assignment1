### Cloud Technologies Assignment 1

- Overview:
    - In this assignment (which is part of your overall portfolio), you are going to develop a REST web application in Go that provides the client with information about books available in a given language based on the Gutenberg library (which holds classic books - most of which are now in the public domain - in a wide range of languages). The service should further determine the number of potential readers (as a second endpoint) presumed to be able to read books in that language. For this purpose, you should interrogate existing web services and return the result in a given output format.
    - The REST web services you will be using for this purpose are:
        - Gutendex API:
            - Endpoint: http://129.241.150.113:8000/books/
            - Documentation: http://129.241.150.113:8000/
        - Language2Countries API:
            - Endpoint: http://129.241.150.113:3000/language2countries/
            - Documentation: http://129.241.150.113:3000/
        - REST Countries API:
            - Endpoint: http://129.241.150.113:8080/v3.1
            - Documentation: http://129.241.150.113:8080/
    - The first API focuses on the provision of Gutenberg library information, whereas the second one provides the mapping from language to country information. The last service provides detailed country information. You will need to use all services in order to complete the assignment.
    - The API documentation is provided under the corresponding links, and all services vary vastly with respect to feature set and quality of documentation. Use Postman or any other REST client to explore the APIs (browser is only of moderate help due to formatting), but be mindful of rate-limiting.
    - General notes on using third-party services in development:
        - When you develop your services that interrogate existing services, try to find the most efficient way of retrieving the necessary information. This generally means reducing the number of requests to these services to a minimum by using the most suitable endpoint that those APIs provide. Consider mocking those services based on exemplary outputs (i.e., develop a simplified version of the third-party service that provides an example response) that you can use to develop your service against locally, before invoking the actual APIs (and causing actual load).
        - While both services are publicly available services, for the course we are only using self-hosted versions (since we might otherwise put unnecessary load on those services). Please ensure you only use the URLs provided above. If they "go down", please let me know. It is better to bring down our own service than the public ones.
        - Directly integrating the data basis of the third-party services (e.g., downloading the underlying raw data) into your service is not permissible, a) since the purpose of the course is to effectively interrogate third-party services that operate as black boxes (in reality you cannot control external services), and b) since that would defeat the value of your service to accommodate dynamic data sources (your point is not to provide data yourself, but to recombine existing and potentially changing data in real time).
    - The final web service should be deployed on Render. The initial development should occur on your local machine and stored in a dedicated workspace repository (more below). However, the actual deployment, you will need to use an additional private Github repository from which Render builds the service. For the submission, you will need to provide both a URL to the deployed Render service as well as your code repository (in the workspace on this Gitlab instance - details below). Again, only turn to this aspect once you have learned about Render. Until then, develop and test locally.
- Specification:
    - The implementation of your service should follow this specification, i.e., the schemas (syntax) of request and response messages, alongside method and status codes should correspond to the ones provided below. Requests and responses may be expressed in exemplary form to illustrate the structure in populated form.
    - Endpoints:
        - Your web service will have three resource root paths:
            - `/librarystats/v1/bookcount/`
            - `/librarystats/v1/readership/`
            - `/librarystats/v1/status/`
        - Assuming your web service should run on localhost, port 8080, your resource root paths would look something like this:
            - `http://localhost:8080/librarystats/v1/bookcount/`
            - `http://localhost:8080/librarystats/v1/readership/`
            - `http://localhost:8080/librarystats/v1/status/`
        - A call to any endpoint should display user-readable guidance on how to invoke the service where necessary (not necessary for the status endpoint, of course). However, both setup and use should be documented for in a Readme (use the repository markdown language) alongside your codebase (more details below).
        - The supported request/response pairs are specified in the following.
        - For the specifications, the following syntax applies:
            - `{:value}` indicates mandatory input parameters specified by the user (i.e., the one using *your* service).
            - `{value}` indicates optional input specified by the user (i.e., the one using *your* service), where `value' can itself contain further optional input.
            - `{value+}` indicates one or more comma-separated optional input.
    - Bookcount Endpoint: Return book count for a given language(s):
        - The initial endpoint returns the count of books for any given language, identified via country 2-letter language ISO codes (ISO 639 Set 1), as well as the number of unique authors. This can be a single as well as multiple languages (comma-separated language codes).
        - Notes:
            - For some books, the author/s is/are unknown. Omit those authors (not the books!) from the count.
            - It goes without saying that not all languages are represented in the database, and sometimes not as fine-grained as it could be (e.g., Norwegian is `no`; there is no differentiation into bokmål (`nb`) and nynorsk (`nn`), even though distinctive language codes exist).
        - Request:
            
            ```
            Method: GET
            Path: bookcount/?language={:two_letter_language_code+}/
            ```
            
            - `two_letter_language_code` is the corresponding [2-letter language ISO codes (ISO 639 Set 1)](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes)
            - Example requests:
                - bookcount/?language=no
                - bookcount/?language=no,fi
        - Response:
            - Content type: `application/json`
            - Status code: 200 if everything is OK, appropriate error code otherwise. Ensure to deal with errors gracefully.
            - Body (Example):
                
                ```
                [
                  {
                     "language": "no",
                     "books": 21,
                     "authors": 14,
                     "fraction": 0.0005
                  },
                  {
                     "language": "fi",
                     "books": 2798,
                     "authors": 228,
                     "fraction": 0.0671
                  }
                ]
                ```
                
            - Notes:
                - The language code is the same as the input code.
                - The book and author information is highlighting the number of available books (from the gutendex) in the queried language.
                - The fraction is the number of books divided by all books served via gutendex.
                - The service should be generic (i.e., allow open entry for any language).
                - The numbers reported in the response are not necessarily the correct answer, but reflect the format that the service should return. (As indicated above, the service should allow for single or multiple languages.)
            - Hint: Think about some known test cases first (’known’ as in ‘you know the results’), before you develop. This way you know the results your service should provide, which reduces the opportunity for bugs in your codebase. Also beware of some oddities that you may only discover during developing or testing the service (a typical challenge of dealing with real-world APIs!).
    - Readership Endpoint: Return the number of potential readers for a given language:
        - The second endpoint should return the number of potential readers for books in a given language, i.e., the population per country in which that language is official (and hence assuming that the inhabitants can potentially read it). This should be reported in addition to the number of books and authors associated with a given language. To this end, use the provided services (inspect them and plan how you can achieve this before engaging in the implementation).
        - Request:
            
            ```
            Method: GET
            Path: readership/{:two_letter_language_code}{?limit={:number}}
            ```
            
            - `{:two_letter_language_code}` refers to the ISO639 Set 1 identifier of the language for which you establish readership.
            - `{?limit={:number}}` is an optional parameter that limits the number of country entries that are reported (in addition to the total number).
            - Example requests:
                - `readership/no`
                - `readership/no/?limit=5`
        - Response:
            - Content type: `application/json`
            - Status code: 200 if everything is OK, appropriate error code otherwise. Ensure to deal with errors gracefully.
            - Body (Example):
                
                ```
                [
                  {
                     "country": "Norway",
                     "isocode": "NO",
                     "books": 21,
                     "authors": 14,
                     "readership": 5379475
                  },
                  {
                     "country": "Svalbard and Jan Mayen",
                     "isocode": "SJ",
                     "books": 21,
                     "authors": 14,
                     "readership": 2562
                  },
                  {
                     "country": "Iceland",
                     "isocode": "IS",
                     "books": 21,
                     "authors": 14,
                     "readership": 366425
                  }
                ]
                ```
                
            - Note:
                - The isocode field should be the two-letter ISO code for the country (recall: country != language) (as per ISO3166-1-Alpha-2).
                - The readership is the `population` of that country.
                - The country name should be the English country name.
                - The book and author information is - as with the previous endpoint - highlighting the number of available books and authors (from gutendex) in the queried language.
                - This endpoint only needs to support input of a single (not multiple) languages at a given time.
    - Diagnostics Endpoint: Getting a status overview of services:
        - The diagnostics interface indicates the availability of individual services this service depends on. The reporting occurs based on status codes returned by the dependent services, and it further provides information about the uptime of the service.
        - Request:
            
            ```
            Method: GET
            Path: status/
            ```
            
        - Response:
            - Content type: `application/json`
            - Status code: 200 if everything is OK, appropriate error code otherwise.
            - Body:
                
                ```
                {
                   "gutendexapi": "<http status code for gutendex API>",
                   "languageapi: "<http status code for language2countries API>",
                   "countriesapi": "<http status code for restcountries API>",
                   "version": "v1",
                   "uptime": <time in seconds from the last service restart>
                }
                ```
                
            - Note: `<some value>` indicates placeholders for values to be populated by the service.
        - **UPDATE**: The values for the different services can either be string (with a content generated based on status code) or integer (with the value being the status code value itself). For this endpoint, the spec is a bit more generous, since it should mostly serve you as a developer.
