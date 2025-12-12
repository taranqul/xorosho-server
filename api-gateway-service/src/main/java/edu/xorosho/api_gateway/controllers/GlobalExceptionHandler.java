package edu.xorosho.api_gateway.controllers;

import edu.xorosho.api_gateway.domains.common.dto.ErrorResponse;
import edu.xorosho.api_gateway.domains.tasks.exeptions.TaskNotFoundExeption;
import edu.xorosho.api_gateway.domains.tasks.exeptions.TaskSchemaNotValid;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import jakarta.servlet.http.HttpServletRequest;


@ControllerAdvice
public class GlobalExceptionHandler {

    @ExceptionHandler(TaskSchemaNotValid.class)
    public ResponseEntity<ErrorResponse> handleInvalidSchema(TaskSchemaNotValid ex, HttpServletRequest request) {
        ErrorResponse error = new ErrorResponse(
                HttpStatus.BAD_REQUEST.value(),
                "Task schema not valid",
                ex.getMessage(),
                request.getRequestURI()
        );
        return new ResponseEntity<>(error, HttpStatus.BAD_REQUEST);
    }

    @ExceptionHandler(TaskNotFoundExeption.class)
    public ResponseEntity<ErrorResponse> handleTaskNotFound(TaskNotFoundExeption ex, HttpServletRequest request) {
        ErrorResponse error = new ErrorResponse(
                HttpStatus.NOT_FOUND.value(),
                "Task not found",
                ex.getMessage(),
                request.getRequestURI()
        );
        return new ResponseEntity<>(error, HttpStatus.NOT_FOUND);
    }

    @ExceptionHandler(Exception.class)
    public ResponseEntity<ErrorResponse> handleOtherExceptions(Exception ex, HttpServletRequest request) {
        ErrorResponse error = new ErrorResponse(
                HttpStatus.INTERNAL_SERVER_ERROR.value(),
                "Internal server error",
                ex.getMessage(),
                request.getRequestURI()
        );
        return new ResponseEntity<>(error, HttpStatus.INTERNAL_SERVER_ERROR);
    }
}