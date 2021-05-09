package ru.thetenzou.tsoddservice.repository.schedule

import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.data.rest.core.annotation.RepositoryRestResource
import ru.thetenzou.tsoddservice.model.schedule.ScheduledTask

@RepositoryRestResource
interface ScheduledTaskRepository : JpaRepository<ScheduledTask, Long>