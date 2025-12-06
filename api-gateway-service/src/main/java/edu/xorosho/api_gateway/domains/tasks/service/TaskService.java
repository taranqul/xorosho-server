package edu.xorosho.api_gateway.domains.tasks.service;

import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

import org.springframework.stereotype.Service;

import edu.xorosho.api_gateway.domains.tasks.dto.TaskRequest;
import edu.xorosho.api_gateway.domains.tasks.dto.TaskResponse;
import edu.xorosho.api_gateway.domains.tasks.models.TaskManagerRequest;
import edu.xorosho.api_gateway.domains.tasks.repositories.UrlRepository;
import lombok.AllArgsConstructor;

@Service
@AllArgsConstructor
public class TaskService {
    private final UrlRepository url_repo;
    private final TaskManager task_manager;
    public TaskResponse createTask(TaskRequest request){
        TaskManagerRequest task_request = new TaskManagerRequest(UUID.randomUUID(), request.getTask_type(), request.getObjects(), request.getPayload());
        UUID uuid = task_manager.createTask(task_request);
        Map<String, String> urls = new HashMap<>();
        for (Map.Entry<String, String> entry : request.getObjects().entrySet()) {
            String object = entry.getKey();
            String type = entry.getValue();
            String name = uuid + "_" + object + type;
            urls.put(object, url_repo.getUrl(name));
        }
        
        return new TaskResponse(uuid.toString(), urls);
    }

    public String getTaskStatus(String id) {
        return task_manager.getTaskStatus(id);
    }
}
