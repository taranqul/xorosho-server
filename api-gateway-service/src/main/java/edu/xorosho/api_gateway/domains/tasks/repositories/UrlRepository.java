package edu.xorosho.api_gateway.domains.tasks.repositories;

import java.net.URI;
import java.net.URLEncoder;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.charset.StandardCharsets;

import org.springframework.stereotype.Repository;

import com.fasterxml.jackson.databind.ObjectMapper;

import lombok.SneakyThrows;

@Repository
public class UrlRepository {
    private final static String baseUrl = "http://storage-gateway-service:8080/storage/external";
    private final static ObjectMapper mapper = new ObjectMapper();
    
    @SneakyThrows
    public String getUploadUrl(String file){
        String query = "filename=" + URLEncoder.encode(file, StandardCharsets.UTF_8)
                     + "&bucketname=" + URLEncoder.encode("upload", StandardCharsets.UTF_8);

        URI uri = new URI(baseUrl + "/uploadUrl" + "?" + query);

        HttpClient client = HttpClient.newHttpClient();

        HttpRequest request = HttpRequest.newBuilder()
                .uri(uri)
                .GET()
                .build();
        try{
            HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
            return mapper.readValue(response.body(), String.class);
        } catch (Exception e) {
            System.out.println(e);
            throw e;
        }

        
    }

    @SneakyThrows
    public String getDownloadUrl(String file){
        String query = "filename=" + URLEncoder.encode(file, StandardCharsets.UTF_8)
                     + "&bucketname=" + URLEncoder.encode("result", StandardCharsets.UTF_8);

        URI uri = new URI(baseUrl + "/downloadUrl"+ "?" + query);

        HttpClient client = HttpClient.newHttpClient();

        HttpRequest request = HttpRequest.newBuilder()
                .uri(uri)
                .GET()
                .build();
        try{
            HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
            return mapper.readValue(response.body(), String.class);
        } catch (Exception e) {
            System.out.println(e);
            throw e;
        }

        
    }
}
