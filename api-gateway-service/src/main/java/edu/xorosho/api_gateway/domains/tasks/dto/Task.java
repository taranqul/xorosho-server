package edu.xorosho.api_gateway.domains.tasks.dto;

import java.util.Map;

import com.fasterxml.jackson.databind.JsonNode;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Task {
    private String id;
    private String type;
    private String status;
    private Map<String, String> objects;
    private JsonNode payload;
}
