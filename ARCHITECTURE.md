# High-level view

At the highest level, the components of a system using Temporal fall into two categories:

- **User-hosted processes**

  - The user's application communicates with the Temporal server using one of the [Temporal SDKs](https://docs.temporal.io/dev-guide).
  - The user runs Temporal [Worker](https://docs.temporal.io/workers) processes. These also use one of the Temporal SDKs and host the user's [Workflow](https://docs.temporal.io/workflows) and [Activity](https://docs.temporal.io/activities) code.

- **Temporal Server**<br>
  Users can host and operate the Temporal server and its database themselves, or use [Temporal Cloud](https://temporal.io/cloud).

<!-- https://lucid.app/lucidchart/0202e4b8-5258-4cd6-a6a0-67159300532b/edit -->
<img width="1521" alt="image" src="https://github.com/temporalio/temporal/assets/52205/0330b3c9-e1eb-4cd2-b9bf-3c6443e56a75">

# Workflow lifecycle

Below we follow a typical sequence of events in the execution of the following very simple workflow:

```
myWorkflow() {
   result = callActivity(myActivity)
   return result
}
```

---

**1. The User Application uses a Temporal SDK to send a `StartWorkflowExecution` request to the Frontend service.**

```mermaid
sequenceDiagram
User Application->>Frontend: StartWorkflowExecution
Frontend->> History: StartWorkflowExecution
History ->> Persistence: CreateWorkflowExecution
Persistence ->> Persistence: Persist MutableState and history tasks
Persistence ->> History: Create Succeed
History->>Frontend: Start Succeed
Frontend->>User Application: Start Succeed
loop QueueProcessor
    History->>Persistence: GetHistoryTasks
		History->>History: ProcessTask
		History->>Matching: AddWorkflowTask
end
```

- The Frontend Service uses a History Service client to call the [`StartWorkflow` handler](https://github.com/temporalio/temporal/blob/ef49189005b5323c532264287af6c08a447aab8a/service/history/api/startworkflow/api.go#L157).
- This initializes history events with a `WorkflowExecutionStarted` event, and persists new mutable state and a history task (transfer task).
- A [queue processor](https://github.com/temporalio/temporal/blob/ef49189005b5323c532264287af6c08a447aab8a/service/history/history_engine.go#L303) goroutine runs for every history task queue.
- The [transfer task queue processor](https://github.com/temporalio/temporal/blob/ef49189005b5323c532264287af6c08a447aab8a/service/history/queues/queue_immediate.go#L150) adds a Workflow Task in the appropriate task queue in the Matching Service.

---

**2. The Worker dequeues the Workflow Task, advances the workflow execution, and becomes blocked on the Activity call.**

```mermaid
sequenceDiagram
Worker->>Frontend: PollWorkflowTask
Frontend->>Matching: PollWorkflowTask
History->>Matching: AddWorkflowTask
Matching->>History: RecordWorkflowTaskStarted
History->>Persistence: UpdateWorkflowExecution
Persistence->>Persistence: Update MutableState & Add timeout timer
Persistence->>History: Update Succeed
History->>Matching: Record Succeed
Matching->>Frontend: WorkflowTask
Frontend->>Persistence: GetHistoryEvents
Persistence->>Frontend: History Events
Frontend->>Worker: WorkflowTask
loop Replayer
    Worker->>Worker: ProcessEvent
end
```

---

**3. The Worker sends a `ScheduleActivityTask`; an Activity task is added in the Matching service.**

```mermaid
sequenceDiagram
Worker ->> Frontend: RespondWorkflowTaskCompleted(ScheduleActivityTask)
Frontend->> History: RespondWorkflowTaskCompleted(ScheduleActivityTask)
History ->> Persistence: UpdateWorkflowExecution
Persistence ->> Persistence: Persist MutableState and history tasks
Persistence ->> History: Update Succeed
History->>Frontend: Respond Succeed
Frontend->>Worker: Respond Succeed
loop QueueProcessor
    History->>Persistence: GetHistoryTasks
		History->>History: ProcessTask
		History->>Matching: AddActivityTask
end
```

---

**4. The Worker dequeues the Activity task**

```mermaid
sequenceDiagram
title: Activity task start
Worker->>Frontend: PollActivityTask
Frontend->>Matching: PollActivityTask
History->>Matching: AddActivityTask
Matching->>History: RecordActivityStarted
History->>Persistence: UpdateWorkflowExecution
Persistence->>Persistence: Update MutableState, add timeout timer
Persistence->>History: Update succeed
History->>Matching: Record Succeed
Matching->>Frontend: ActivityTask
Frontend->>Worker: ActivityTask
```

---

**4. The Worker sends `RespondActivityCompleted` to the History service; a Workflow Task is added to the Matching service**

```mermaid
sequenceDiagram
SDK->>Frontend: RespondActivityCompleted
Frontend->>History: RespondActivityCompleted
History->>Persistence: UpdateWorkflowExecution
Persistence->>Persistence: Update MutableState & add transfer task
Persistence->>History: Update Succeed
History->>Frontend: Respond Succeed
Frontend->>SDK: Respond Succeed
loop QueueProcessor
    History->>Persistence: GetHistoryTasks
		History->>History: ProcessTask
		History->>Matching: AddWorkflowTask
end
```

---

**5. The Worker dequeues the Workflow Task, advances the workflow, and finds that it has reached its end**

\<Same sequence diagram as step 2 above\>

---

**6. The Worker sends `RespondWorkflowTaskCompleted` to the History Service**

```mermaid
sequenceDiagram
SDK->>Frontend: RespondWorkflowTaskCompleted
Frontend->>History: RespondWorkflowTaskCompleted
History->>Persistence: UpdateWorkflowExecution
Persistence->>Persistence: Update MutableState & add tasks (visibility, tiered storage, retention etc)
Persistence->>History: Update Succeed
History->>Frontend: Respond Succeed
Frontend->>SDK: Respond Succeed
loop QueueProcessor
    History->>Persistence: GetHistoryTasks
		History->>History: ProcessTask (Update visibility, Upload to S3, Delete data etc)
end
```
