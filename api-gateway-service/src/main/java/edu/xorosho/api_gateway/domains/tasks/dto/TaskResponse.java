package edu.xorosho.api_gateway.domains.tasks.dto;

import java.util.Map;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class TaskResponse {
   private String taskId;
   private Map<String, String> urls;
}