package edu.xorosho.api_gateway.integration;

import com.fasterxml.jackson.databind.ObjectMapper;
import edu.xorosho.api_gateway.domains.tasks.dto.TaskRequest;
import edu.xorosho.api_gateway.domains.tasks.models.TaskSchema;
import edu.xorosho.api_gateway.domains.tasks.service.TaskSchemeValidator;
import edu.xorosho.api_gateway.domains.tasks.service.TaskService;
import edu.xorosho.api_gateway.domains.tasks.exeptions.TaskNotFoundExeption;
import edu.xorosho.api_gateway.domains.tasks.exeptions.TaskSchemaNotValid;
import edu.xorosho.api_gateway.domains.status.service.StatusService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.bean.override.mockito.MockitoBean;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;

import java.util.Map;
import java.util.List;

import static org.mockito.Mockito.*;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

@SpringBootTest
@AutoConfigureMockMvc
class ApiGatewayIntegrationTest {

    @Autowired
    private MockMvc mockMvc;

    @Autowired
    private ObjectMapper objectMapper;

    @MockitoBean
    private TaskService taskService;

    @MockitoBean
    private TaskSchemeValidator taskSchemeValidator;

    @MockitoBean
    private StatusService statusService;

    @Test
    void postTask_validRequest_returnsOk() throws Exception {
        TaskRequest request = new TaskRequest("test", Map.of("file", ".mp4"), null);

        mockMvc.perform(post("/api/task")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(request)))
                .andExpect(status().isOk());

        verify(taskSchemeValidator).validateScheme(any(TaskRequest.class));
        verify(taskService).createTask(any(TaskRequest.class));
    }

    @Test
    void postTask_invalidSchema_returnsBadRequest() throws Exception {
        TaskRequest request = new TaskRequest("invalid", Map.of(), null);

        doThrow(new TaskSchemaNotValid()).when(taskSchemeValidator).validateScheme(any(TaskRequest.class));

        mockMvc.perform(post("/api/task")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(request)))
                .andExpect(status().isBadRequest());
    }

    @Test
    void getTaskSchemas_existingTask_returnsOk() throws Exception {
        String taskName = "existing";
        when(taskSchemeValidator.getTaskSchema(taskName)).thenReturn(new TaskSchema(null));

        mockMvc.perform(get("/api/task/{task}/schema", taskName))
                .andExpect(status().isOk());
    }

    @Test
    void getTaskSchemas_notFound_returnsNotFound() throws Exception {
        String taskName = "missing";
        when(taskSchemeValidator.getTaskSchema(taskName)).thenThrow(new TaskNotFoundExeption());

        mockMvc.perform(get("/api/task/{task}/schema", taskName))
                .andExpect(status().isNotFound());
    }

    @Test
    void getTasks_returnsList() throws Exception {
        when(taskSchemeValidator.getTasks()).thenReturn(List.of("task1", "task2"));

        mockMvc.perform(get("/api/task"))
                .andExpect(status().isOk())
                .andExpect(content().json("[\"task1\",\"task2\"]"));
    }

    @Test
    void ping_returnsPong() throws Exception {
        mockMvc.perform(get("/api/status/ping"))
                .andExpect(status().isOk())
                .andExpect(content().string("pong"));
    }

    @Test
    void getSummary_returnsStatusOk() throws Exception {
        when(statusService.getSummary()).thenReturn(new edu.xorosho.api_gateway.domains.status.dto.StatusResponse("healthy"));
        
        mockMvc.perform(get("/api/status/summary"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.status").value("healthy"));
    }
}