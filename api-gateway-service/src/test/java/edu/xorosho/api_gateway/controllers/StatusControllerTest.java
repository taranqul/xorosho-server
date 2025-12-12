package edu.xorosho.api_gateway.controllers;

import edu.xorosho.api_gateway.domains.status.dto.StatusResponse;
import edu.xorosho.api_gateway.domains.status.service.StatusService;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.http.MediaType;
import org.springframework.test.context.bean.override.mockito.MockitoBean;
import org.springframework.test.web.servlet.MockMvc;

import static org.mockito.Mockito.verify;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(StatusController.class)
class StatusControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockitoBean
    private StatusService statusService;

    @Test
    void ping_returnsPong() throws Exception {
        mockMvc.perform(get("/api/status/ping"))
                .andExpect(status().isOk())
                .andExpect(content().string("pong"));
    }

    @Test
    void getSummary_returnsStatusResponse() throws Exception {
        StatusResponse response = new StatusResponse("healthy");
        Mockito.when(statusService.getSummary()).thenReturn(response);

        mockMvc.perform(get("/api/status/summary")
                        .contentType(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(content().json("{\"status\":\"healthy\"}"));

        verify(statusService).getSummary();
    }
}