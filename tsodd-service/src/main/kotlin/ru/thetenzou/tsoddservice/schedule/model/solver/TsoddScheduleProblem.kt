package ru.thetenzou.tsoddservice.schedule.model.solver

import org.optaplanner.core.api.domain.solution.PlanningEntityCollectionProperty
import org.optaplanner.core.api.domain.solution.PlanningScore
import org.optaplanner.core.api.domain.solution.PlanningSolution
import org.optaplanner.core.api.domain.solution.ProblemFactProperty
import org.optaplanner.core.api.domain.valuerange.ValueRangeProvider
import org.optaplanner.core.api.score.buildin.hardmediumsoft.HardMediumSoftScore
import ru.thetenzou.tsoddservice.crew.model.Crew
import java.time.LocalDate

/**
 * TsoddScheduleProblem planning solution for OptaPlanner * holds data required for OptaPlanner to solve problem * and solution of the problem
 *
 * @param scheduleId id of the schedule in database
 * @param availableDates list of available dates for planning
 * @param availableCrews list of crew available for assigning tasks
 *
 * @param planningTaskList list of all tasks that has to be planned
 */
@PlanningSolution
class TsoddScheduleProblem(

    val scheduleId: Long?,

    @ProblemFactProperty
    val parameters: ScheduleParameters?,

    @ValueRangeProvider(id = "dateRange")
    val availableDates: List<LocalDate>?,

    @ValueRangeProvider(id = "crewRange")
    val availableCrews: List<Crew>?,

    @PlanningEntityCollectionProperty
    val planningTaskList: List<PlannedTask>?,
) {
    @PlanningScore
    lateinit var score: HardMediumSoftScore

    /**
     * default constructor required for OptaPlanner
     */
    constructor() : this(null, null, null, null, null)
}

data class ScheduleParameters (
    val resourceLimit: Double,
)