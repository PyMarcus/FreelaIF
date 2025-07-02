package com.freelaif.api_gateway;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.cloud.context.config.annotation.RefreshScope;

@RefreshScope
@RestController
public class TesteController {

    @Value("${message.test}")
    private String minhaVariavel;

    @GetMapping("/variavel")
    public String imprimirVariavel() {
        return "Valor da variável: " + minhaVariavel;
    }
}
