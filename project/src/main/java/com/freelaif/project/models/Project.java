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
@Document(collation = "projects")
public class Project {
    @Id
    private int id;
    private int clientId;
    private Integer developerId;
    private String title;
    private String description;
    private LocalDate createdAt;
    private LocalDate dateLimit;
    public enum Status {
        CREATED,
        COMPLETE,
        PROGRESS,
        CANCELED
    }
}
