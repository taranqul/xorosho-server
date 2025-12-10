package edu.xorosho.api_gateway.controllers;

import edu.xorosho.api_gateway.domains.tasks.dto.TaskRequest;
import edu.xorosho.api_gateway.domains.tasks.dto.TaskResponse;
import edu.xorosho.api_gateway.domains.tasks.dto.TaskSchemaResponse;
import edu.xorosho.api_gateway.domains.tasks.service.TaskService;
import edu.xorosho.api_gateway.domains.tasks.service.TaskSchemeValidator;
import io.swagger.v3.oas.annotations.tags.Tag;
import lombok.AllArgsConstructor;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@Tag(name = "Task-related endpoints")
@RestController
@AllArgsConstructor
@RequestMapping("/api/task")
public class TaskController {

    private final TaskService taskService;
    private final TaskSchemeValidator taskSchemeValidator;

    @PostMapping
    public TaskResponse postTask(@RequestBody TaskRequest request) {
        return taskService.createTask(request);
    }

    @GetMapping("/{task}/schema")
    public TaskSchemaResponse getTaskSchemas(@PathVariable String task) {
        return new TaskSchemaResponse(taskSchemeValidator.getTaskSchema(task).getSchema());
    }

    @GetMapping("/{task}/status")
    public String getTaskStatus(@PathVariable String task) {
        return taskService.getTaskStatus(task);
    }

    @GetMapping
    public List<String> getTasks() {
        return taskSchemeValidator.getTasks();
    }
}