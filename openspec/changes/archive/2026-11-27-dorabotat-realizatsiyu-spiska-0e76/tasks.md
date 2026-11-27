## 1. Сброс поиска без мигания при закрытии

- [x] 1.1 В `front/src/App.vue`, watcher `projectsMenu`: убрать `this.projectSearch = ''` из ветки закрытия (`if (!v) return;`) и выполнять сброс в начале ветки открытия, перед планированием фокуса
- [x] 1.2 Убедиться, что `onProjectSearchKeydown`, `filteredProjects`, `normalizedProjectSearch` и логика фокуса (`requestAnimationFrame`/`setTimeout`) не изменены
- [x] 1.3 Проверить, что watcher `filteredProjects` — одна строка `this.$nextTick(this.$refs.projectsMenu?.updateDimensions);`, обёртка поля — `class="pa-2"`, правила `.project-search` в `<style>` нет

## 2. Проверка

- [x] 2.1 Тесты: автотестов фронтенда в репозитории нет, новых Go-тестов задача не требует; убедиться, что `go test ./...` не затронут (Go-код не менялся), и прогнать `cd front && npm run lint`
- [x] 2.2 Ручная проверка в браузере: ввести запрос, кликнуть найденный проект — при закрытии не мелькает полный список
- [x] 2.3 Ручная проверка: повторно открыть меню — поле пустое, полный список, фокус в поле
- [x] 2.4 Ручная проверка: запрос, оставляющий 1–2 проекта, — меню прижато к кнопке «Project»; закрыть и открыть — меню на обычном месте
- [x] 2.5 Ручная проверка: Escape закрывает меню, Ctrl/Cmd+K открывает «Go to…», «no matching projects» и «new project» отображаются как раньше

## 3. Git

- [ ] 3.1 Не добавлять в индекс `openspec/*`, `run.json`, `.gitkeep`, `front/front_test.go`; коммит содержит только `front/src/App.vue`
- [ ] 3.2 Проверить `git diff 6927e8f --stat`: только `front/src/App.vue`; если промежуточные коммиты ветки содержат `openspec/*`, подготовить чистую историю (один коммит поверх `6927e8f` с `front/src/App.vue`), не делая force-push без подтверждения человека
