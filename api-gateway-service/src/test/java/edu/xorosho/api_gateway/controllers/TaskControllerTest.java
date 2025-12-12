package edu.xorosho.api_gateway.controllers;

import com.fasterxml.jackson.databind.ObjectMapper;
import edu.xorosho.api_gateway.domains.tasks.dto.TaskRequest;
import edu.xorosho.api_gateway.domains.tasks.dto.TaskResponse;
import edu.xorosho.api_gateway.domains.tasks.models.TaskSchema;
import edu.xorosho.api_gateway.domains.tasks.service.TaskSchemeValidator;
import edu.xorosho.api_gateway.domains.tasks.service.TaskService;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.http.MediaType;
import org.springframework.test.context.bean.override.mockito.MockitoBean;
import org.springframework.test.web.servlet.MockMvc;

import java.util.List;
import java.util.Map;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.eq;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(TaskController.class)
class TaskControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockitoBean
    private TaskService taskService;

    @MockitoBean
    private TaskSchemeValidator taskSchemeValidator;

    @Autowired
    private ObjectMapper objectMapper;

    @Test
    void postTask_validRequest_callsService() throws Exception {
        TaskRequest request = new TaskRequest("test", Map.of("file", ".mp4"), null);
        TaskResponse response = new TaskResponse("1234", Map.of("file", "http://example.com/1234_file.mp4"));

        Mockito.when(taskService.createTask(any(TaskRequest.class)))
                .thenReturn(response);

        mockMvc.perform(post("/api/task")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(request)))
                .andExpect(status().isOk())
                .andExpect(content().json(objectMapper.writeValueAsString(response)));

        Mockito.verify(taskService).createTask(any(TaskRequest.class));
    }

    @Test
    void getTaskSchemas_callsValidatorAndReturnsOk() throws Exception {
        String taskName = "test-task";
        TaskSchema taskSchema = new TaskSchema(null);

        Mockito.when(taskSchemeValidator.getTaskSchema(taskName))
                .thenReturn(taskSchema);

        mockMvc.perform(get("/api/task/{task}/schema", taskName))
                .andExpect(status().isOk());

        Mockito.verify(taskSchemeValidator).getTaskSchema(eq(taskName));
    }

    @Test
    void getTaskStatus_returnsStatusString() throws Exception {
        String taskName = "test-task";
        String statusStr = "IN_PROGRESS";

        Mockito.when(taskService.getTaskStatus(taskName))
                .thenReturn(statusStr);

        mockMvc.perform(get("/api/task/{task}/status", taskName))
                .andExpect(status().isOk())
                .andExpect(content().string(statusStr));

        Mockito.verify(taskService).getTaskStatus(eq(taskName));
    }

    @Test
    void getTasks_returnsTasksList() throws Exception {
        List<String> tasks = List.of("task1", "task2");

        Mockito.when(taskSchemeValidator.getTasks())
                .thenReturn(tasks);

        mockMvc.perform(get("/api/task"))
                .andExpect(status().isOk())
                .andExpect(content().json("[\"task1\",\"task2\"]"));

        Mockito.verify(taskSchemeValidator).getTasks();
    }
}