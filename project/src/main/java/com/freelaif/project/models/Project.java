package com.freelaif.project.models;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import java.time.LocalDate;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
@Document(collection = "projects")
public class Project {
    @Id
    private int id;
    private String title;
    private String description;
    private Integer developerId;
    private int clientId;
    private LocalDate createdAt;
    private LocalDate dateLimit;
    private Status status;
    public enum Status {
        CREATED,
        COMPLETE,
        PROGRESS,
        CANCELED
    }
}
