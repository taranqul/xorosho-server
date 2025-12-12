package edu.xorosho.api_gateway.domains.tasks.dto;

import com.fasterxml.jackson.databind.JsonNode;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.Map;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class TaskRequest {
    private String taskType;
    private Map<String, String> objects;
    private JsonNode payload;
}