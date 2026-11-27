## 1. Воспроизведение
- [x] 1.1 Написан падающий тест `front/front_test.go` (`TestProjectsMenuRepositionsOnSearch`): проверяет, что у `v-menu` проектов есть `ref` и что при изменении результатов поиска вызывается `updateDimensions`/`onResize`; `go test ./front/` красный

## 2. Исправление
- [x] 2.1 В `front/src/App.vue` добавить `ref="projectsMenu"` на `<v-menu v-if="user" v-model="projectsMenu" ...>`
- [x] 2.2 В `watch` добавить `filteredProjects()`: в `this.$nextTick` получить `this.$refs.projectsMenu` и, если он есть и `isActive`, вызвать `updateDimensions()`
- [ ] 2.3 Прогнать `npm run lint` и `npm run fmt-check` в `front/`

## 3. Проверка
- [ ] 3.1 Тесты: `go test ./front/` зелёный, `go test ./...` зелёный
- [ ] 3.2 Вручную: ~20+ проектов, открыть меню «Project», ввести «test1» — меню опускается и прижато к кнопке «Project»; очистить запрос — меню снова поднимается и не выходит за экран
- [ ] 3.3 Вручную: на маленьком окне меню не выходит за границы экрана; фокус остаётся в поле поиска при вводе
