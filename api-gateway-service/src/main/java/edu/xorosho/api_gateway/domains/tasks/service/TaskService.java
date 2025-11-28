package edu.xorosho.api_gateway.domains.tasks.service;

import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

import org.springframework.stereotype.Service;

import edu.xorosho.api_gateway.domains.tasks.dto.TaskRequest;
import edu.xorosho.api_gateway.domains.tasks.dto.TaskResponse;
import edu.xorosho.api_gateway.domains.tasks.repositories.UrlRepository;
import lombok.AllArgsConstructor;

@Service
@AllArgsConstructor
public class TaskService {
    private final UrlRepository url_repo;
    public TaskResponse createTask(TaskRequest request){
        UUID uuid = UUID.randomUUID();
        Map<String, String> urls = new HashMap<>();
        for (Map.Entry<String, String> entry : request.getObjects().entrySet()) {
            String object = entry.getKey();
            String type = entry.getValue();
            String name = uuid + "_" + object + type;
            urls.put(object, url_repo.getUrl(name));
        }
        
        return new TaskResponse(uuid.toString(), urls);
    }
}
