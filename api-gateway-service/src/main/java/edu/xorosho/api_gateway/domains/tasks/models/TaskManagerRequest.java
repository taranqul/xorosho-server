package edu.xorosho.api_gateway.domains.tasks.models;

import java.util.Map;
import java.util.UUID;

import com.fasterxml.jackson.databind.JsonNode;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class TaskManagerRequest {
    public UUID Id;
    private String Type;
    private Map<String, String> Objects;
    private JsonNode Payload;
}
