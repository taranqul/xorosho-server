package edu.xorosho.api_gateway.domains.tasks.service;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.UUID;

import org.springframework.stereotype.Service;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import lombok.AllArgsConstructor;
import lombok.SneakyThrows;
import edu.xorosho.api_gateway.domains.tasks.models.TaskManagerRequest;

@Service
@AllArgsConstructor
public class TaskManager {
    private final ObjectMapper mapper = new ObjectMapper();
    private final HttpClient client = HttpClient.newHttpClient();
    @SneakyThrows
    public UUID createTask(TaskManagerRequest task) {
        String body = mapper.writeValueAsString(task);
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create("http://task-manager-service:8080/task"))
                .header("Content-Type", "application/json")
                .POST(HttpRequest.BodyPublishers.ofString(body))
                .build();

        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
        String responseBody = response.body();
        JsonNode root = mapper.readTree(responseBody);
        JsonNode valueNode = root.get("uuid");

        if (valueNode == null) {
            throw new RuntimeException("cant read json, empty uuid");
        }

        String value = valueNode.asText();

        return UUID.fromString(value);
    }
}
