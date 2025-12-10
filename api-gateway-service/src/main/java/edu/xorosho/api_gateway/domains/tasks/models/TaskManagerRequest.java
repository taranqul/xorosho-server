package edu.xorosho.api_gateway.domains.tasks.models;

import java.util.Map;
import java.util.UUID;

import com.fasterxml.jackson.databind.JsonNode;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class TaskManagerRequest {
    private UUID id;
    private String type;
    private Map<String, String> objects;
    private JsonNode payload;
}