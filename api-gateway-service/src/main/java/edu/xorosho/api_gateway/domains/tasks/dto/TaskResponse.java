package edu.xorosho.api_gateway.domains.tasks.dto;

import java.util.Map;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class TaskResponse {
   private String task_id;
   private Map<String, Object> objects;
}
