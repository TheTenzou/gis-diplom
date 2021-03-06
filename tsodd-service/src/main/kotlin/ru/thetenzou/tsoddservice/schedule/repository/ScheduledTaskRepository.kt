package ru.thetenzou.tsoddservice.schedule.repository

import org.springframework.data.domain.Page
import org.springframework.data.domain.Pageable
import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.data.rest.core.annotation.RepositoryRestResource
import org.springframework.data.rest.core.annotation.RestResource
import ru.thetenzou.tsoddservice.schedule.model.Schedule
import ru.thetenzou.tsoddservice.schedule.model.ScheduledTask

@RepositoryRestResource
interface ScheduledTaskRepository : JpaRepository<ScheduledTask, Long> {

    fun getBySchedule(schedule: Schedule, pageable: Pageable): Page<ScheduledTask>

    @RestResource(exported = false)
    fun getBySchedule(schedule: Schedule): List<ScheduledTask>
}