package edu.xorosho.api_gateway.domains.tasks.service;
import com.fasterxml.jackson.core.type.TypeReference;
import org.springframework.stereotype.Service;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.net.URI;
import java.net.URLEncoder;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.charset.StandardCharsets;
import java.util.List;

import edu.xorosho.api_gateway.domains.tasks.dto.TaskRequest;
import lombok.AllArgsConstructor;

@Service
@AllArgsConstructor
public class TaskSchemeValidator {

    private final ObjectMapper objectMapper;
    private final HttpClient client = HttpClient.newHttpClient();
    public void validateScheme(TaskRequest request) throws Exception {

        String body = objectMapper.writeValueAsString(request);

        HttpRequest httpRequest = HttpRequest.newBuilder()
                .uri(URI.create("http://worker-manager-service:8080/worker/validate"))
                .header("Content-Type", "application/json")
                .POST(HttpRequest.BodyPublishers.ofString(body))
                .build();

        HttpResponse<String> response = client.send(httpRequest, HttpResponse.BodyHandlers.ofString());

        if (response.statusCode() != 200) {
            throw new RuntimeException("Validation failed with status: " + response.statusCode());
        }
    }

    public List<String> getTasks() throws Exception {
        URI uri = URI.create("http://worker-manager-service:8080/worker");

        HttpRequest request = HttpRequest.newBuilder()
                .uri(uri)
                .GET()
                .header("Accept", "application/json")
                .build();

        HttpResponse<String> response =
                client.send(request, HttpResponse.BodyHandlers.ofString());

        if (response.statusCode() != 200) {
            throw new RuntimeException(
                    "Request failed, status=" + response.statusCode()
            );
        }

        return objectMapper.readValue(
                response.body(),
                new TypeReference<List<String>>() {}
        );
    }

    public JsonNode getSchema(String name) throws Exception {
        String encodedName = URLEncoder.encode(name, StandardCharsets.UTF_8);

        URI uri = URI.create(
                "http://worker-manager-service:8080/task/scheme?name=" + encodedName
        );

        HttpRequest request = HttpRequest.newBuilder()
                .uri(uri)
                .GET()
                .header("Accept", "application/json")
                .build();

        HttpResponse<String> response =
                client.send(request, HttpResponse.BodyHandlers.ofString());

        if (response.statusCode() != 200) {
            throw new RuntimeException(
                    "Request failed, status=" + response.statusCode()
            );
        }

        JsonNode jsonNode = objectMapper.readTree(response.body());

        return jsonNode;
    }
}
