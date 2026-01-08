package edu.xorosho.api_gateway.domains.tasks.dto;

import lombok.AllArgsConstructor;
import lombok.Data;

import java.util.Map;

import com.fasterxml.jackson.databind.JsonNode;

@Data
@AllArgsConstructor
public class TaskResult {
    private String task_id;
    private Map<String, Object> objects;
    private JsonNode payload;
}
