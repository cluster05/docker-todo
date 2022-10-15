<template>
    <div class="w-screen h-screen flex justify-center items-center bg-gradient-to-r from-fuchsia-300 to-purple-300 ">
        <div class="w-96 rounded-lg bg-white">
            <div class="p-4 border-b-2 flex items-center">
                <img src="@/assets/menu.svg" class="w-5 h-5" alt="menu">
                <h2 class="m-2 ml-4 text-lg text-gray-700 font-bold">Todo List</h2>
            </div>
            <div class="p-4 flex items-center">
                <input type="text" name="title" id="title" placeholder="add new task" v-model="title"
                    @keydown.enter="createTodo" class="w-full px-4 py-2 rounded-md border focus:outline-none">
                <button class="ml-4 p-2 bg-orange-300 rounded" @click="createTodo">
                    <img src="@/assets/close.svg" class="w-5 h-5 rotate-45" alt="menu">
                </button>
            </div>
            <div class="p-4">
                <template v-if="todos.length > 0">
                    <div class="p-3 flex items-center justify-between border-b cursor-default"
                        v-for="(todo,index) in todos" :key="todo.todoId">
                        <div class="flex items-center">
                            <input type="checkbox" class="m-1 w-4 h-4 cursor-pointer" :name="`checkbox_${todo.todoId}`"
                                @click="updateTodo(todo.todoId,index,!todo.isCompleted)" v-model="todo.isCompleted"
                                :id="`checkbox_${todo.todoId}`">
                            <span class="ml-2 text-gray-500"
                                :class="{'line-through':todo.isCompleted}">{{todo.title}}</span>
                        </div>
                        <span class="cursor-pointer" @click="deleteTodo(todo.todoId,index)">
                            <img src="@/assets/close.svg" class="w-5 h-5 " alt="delete">
                        </span>
                    </div>
                </template>
                <template v-else>
                    <div class="py-10 flex justify-center">
                        <span class="font-semibold text-xl text-gray-300 select-none">No task avaliable. Add new
                            task</span>
                    </div>
                </template>
            </div>

        </div>
    </div>
    <div class="absolute right-0 bottom-0">
        <div class="px-4 py-1 text-sm  bg-white cursor-default">
           created by : <a target="_blank" class="text-blue-400 hover:text-blue-600" href="https://github.com/cluster05">Ajay Kumbhar</a>
        </div>
    </div>
</template>

<script>
import { todoservice } from "@/services/todo.service"

export default {
    name: "App",
    mounted() {
        this.readAllTodo()
    },
    data() {
        return {
            title: "",
            todos: [],
        }
    },
    methods: {
        async createTodo() {
            try {
                if (this.title.trim() == "") {
                    this.title = ""
                    return
                }
                const payload = {
                    title: this.title
                }

                const response = await todoservice.post(``, payload)

                this.todos = [response.data.result, ...this.todos]

                this.title = ""
            } catch (error) {
                console.log(error)
            }
        },
        async readAllTodo() {
            try {
                const response = await todoservice.get(``)
                this.todos = response.data.result
            } catch (error) {
                console.log(error)
            }
        },
        async updateTodo(todoId, index, isCompleted) {
            try {

                const payload = {
                    isCompleted: isCompleted
                }
                const response = await todoservice.patch(`/${todoId}`, payload)

                this.todos.splice(index, 1, response.data.result)

            } catch (error) {
                console.log(error)
            }
        },
        async deleteTodo(todoId, index) {
            try {
                await todoservice.delete(`/${todoId}`)
                this.todos.splice(index, 1)
            } catch (error) {
                console.log(error)
            }
        }

    }
}
</script>

<style>

</style>