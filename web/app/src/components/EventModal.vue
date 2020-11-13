<template>
	<div>
		<b-modal
			id="event-modal"
			hide-backdrop
			content-class="shadow"
			title="Detalhes do evento"
			@hide="resetModal"
			@show="loadDataIntoForm"
		>

			<b-form @submit.stop.prevent="submitEvent">
				<b-form-group
					id="title-group"
					label="Título"
					label-for="title"
				>
					<b-form-input
						id="title"
						v-model="$v.form.title.$model"
						:state="validateState('title')"
						placeholder="Nome curto do evento..."
						required
					></b-form-input>
					<b-form-invalid-feedback id="title-live-feedback">Campo obrigatório.</b-form-invalid-feedback>
				</b-form-group>

				<b-form-group
					id="desc-group"
					label="Descrição"
					label-for="desc"
				>
					<b-form-textarea
						id="desc"
						v-model="$v.form.desc.$model"
						:state="validateState('desc')"
						placeholder="Descreva o evento..."
						rows="3"
						max-rows="6"
						required
					></b-form-textarea>
					<b-form-invalid-feedback id="desc-live-feedback">Campo obrigatório.</b-form-invalid-feedback>
				</b-form-group>

				<!-- <b-form-group
					id="title-group"
					label="Título"
					label-for="title"
				>
					<b-form-datepicker
						id="example-datepicker"
						v-model="$v.form.day.$model"
						:state="validateState('day')"
						class="mb-2"
						locale="pt-BR"
						required
					></b-form-datepicker>
					<b-form-invalid-feedback id="title-live-feedback">Campo obrigatório.</b-form-invalid-feedback>
				</b-form-group> -->

				<b-form-group
					id="start-group"
					label="Horário de início"
					label-for="start"
				>
					<b-form-timepicker
						id="start"
						v-model="$v.form.start.$model"
						:state="validateState('start')"
						placeholder="Nenhum horário escolhido"
						locale="pt-BR"
						required
					></b-form-timepicker>
					<b-form-invalid-feedback id="start-live-feedback">Campo obrigatório.</b-form-invalid-feedback>
				</b-form-group>

				<b-form-group
					id="end-group"
					label="Horário de término"
					label-for="end"
				>
					<b-form-timepicker
						id="end"
						v-model="$v.form.end.$model"
						:state="validateState('end')"
						placeholder="Nenhum horário escolhido"
						locale="pt-BR"
						required
					></b-form-timepicker>
					<b-form-invalid-feedback id="end-live-feedback">Campo obrigatório.</b-form-invalid-feedback>
				</b-form-group>

			</b-form>

			<template #modal-footer>
				<div
					v-if="Object.keys(currentEvent).length > 0"
					class="w-100"
				>
					<b-button
						variant="danger"
						size="sm"
						@click="deleteEvent"
					>
						Excluir
					</b-button>
					<b-button
						variant="primary"
						size="sm"
						class="float-right"
						@click="handleUpdate"
					>
						Atualizar
					</b-button>
				</div>

				<div
					v-else
					class="w-100"
				>
					<b-button
						variant="primary"
						size="sm"
						class="float-right"
						@click="handleCreate"
					>
						Criar
					</b-button>
				</div>
			</template>

		</b-modal>
	</div>
</template>

<script>
import { mapActions, mapGetters, mapState } from "vuex";
import { required } from "vuelidate/lib/validators";

export default {
	data() {
		return {
			form: {
				title: "",
				desc: "",
				// day: "",
				start: "",
				end: "",
			},
			submitStatus: null,
		};
	},
	validations: {
		form: {
			title: {
				required,
			},
			desc: {
				required,
			},
			// day: {
			// 	required,
			// },
			start: {
				required,
			},
			end: {
				required,
			},
		},
	},
	methods: {
		...mapActions(["addEvent", "removeEvent"]),
		validateState(name) {
			const { $dirty, $error } = this.$v.form[name];
			return $dirty ? !$error : null;
		},
		submitEvent() {
			this.$v.$touch();
			if (!this.$v.form.$anyError) {
				this.submitStatus = 1;

				var splitStart = this.form.start.split(':')
				var paddedStart = `${splitStart[0].toString().padStart(2, '0')}:${splitStart[1].toString().padStart(2, '0')}:${splitStart[2].toString().padStart(2, '0')}`;
				var newStart = `${this.currentDate.getFullYear()}-${(this.currentDate.getMonth() + 1)
					.toString()
					.padStart(2, "0")}-${this.currentDate
					.getDate()
					.toString()
					.padStart(2, "0")}T${paddedStart}`;

				var splitEnd = this.form.end.split(':')
				var paddedEnd = `${splitEnd[0].toString().padStart(2, '0')}:${splitEnd[1].toString().padStart(2, '0')}:${splitEnd[2].toString().padStart(2, '0')}`;
				var newEnd = `${this.currentDate.getFullYear()}-${(this.currentDate.getMonth() + 1)
					.toString()
					.padStart(2, "0")}-${this.currentDate
					.getDate()
					.toString()
					.padStart(2, "0")}T${paddedEnd}`;

				let payload = {
					title: this.form.title,
					desc: this.form.desc,
					start: new Date(newStart).getTime() / 1000,
					end: new Date(newEnd).getTime() / 1000,
				};

				fetch("/backend/API/event/create", {
					method: "POST",
					headers: {
						"Content-Type": "application/json",
						Authorization: "Bearer: " + this.token,
						_id: JSON.stringify(this.userInfo.ID),
					},
					body: JSON.stringify(payload),
				})
					.then((response) => {
						if (response.status === 200) {
							return response.json();
						} else {
							return "Erro criando evento.";
						}
					})
					.then((response) => {
						if (response === "Erro criando evento.") {
							console.log(response);
						} else {
							payload["_id"] = response;
							this.addEvent(payload);
						}
					});

				// Hide the modal manually
				this.$nextTick(() => {
					this.$bvModal.hide("event-modal");
				});
			}
		},
		resetModal() {
			this.form = {
				title: "",
				desc: "",
				// day: "",
				start: "",
				end: "",
			};
			this.$nextTick(() => {
				this.$v.$reset();
			});
		},
		loadDataIntoForm() {
			if (Object.keys(this.currentEvent).length > 0) {
				var time = new Date(this.currentEvent.start * 1000);
				var start = `${time.getHours()}:${time.getMinutes()}:${time.getSeconds()}`;
				time = new Date(this.currentEvent.end * 1000);
				var end = `${time.getHours()}:${time.getMinutes()}:${time.getSeconds()}`;

				this.form.title = this.currentEvent.title;
				this.form.desc = this.currentEvent.desc;
				this.form.start = start;
				this.form.end = end;
			}
		},
		handleCreate(bvModalEvt) {
			// Prevent modal from closing
			bvModalEvt.preventDefault();
			// Trigger submit handler
			this.submitEvent();
		},
		handleUpdate(bvModalEvt) {
			// Prevent modal from closing
			bvModalEvt.preventDefault();
			// Trigger submit handler
			this.updateEvent();
		},
		deleteEvent() {
			// this.submitStatus = 1;

			fetch("/backend/API/event/delete", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
					Authorization: "Bearer: " + this.token,
					_id: JSON.stringify(this.currentEvent._id),
				},
			})
				.then((response) => {
					if (response.status === 200) {
						return response.json();
					} else {
						return "Erro deletando evento.";
					}
				})
				.then((response) => {
					if (response === "Erro deletando evento.") {
						console.log(response);
					} else {
						this.removeEvent(this.currentEvent._id);
					}
				});

			// Hide the modal manually
			this.$nextTick(() => {
				this.$bvModal.hide("event-modal");
			});
		},
		updateEvent() {
			this.$v.$touch();
			if (!this.$v.form.$anyError) {
				this.submitStatus = 1;

				var splitStart = this.form.start.split(':')
				var paddedStart = `${splitStart[0].toString().padStart(2, '0')}:${splitStart[1].toString().padStart(2, '0')}:${splitStart[2].toString().padStart(2, '0')}`;
				var newStart = `${this.currentDate.getFullYear()}-${(this.currentDate.getMonth() + 1)
					.toString()
					.padStart(2, "0")}-${this.currentDate
					.getDate()
					.toString()
					.padStart(2, "0")}T${paddedStart}`;

				var splitEnd = this.form.end.split(':')
				var paddedEnd = `${splitEnd[0].toString().padStart(2, '0')}:${splitEnd[1].toString().padStart(2, '0')}:${splitEnd[2].toString().padStart(2, '0')}`;
				var newEnd = `${this.currentDate.getFullYear()}-${(this.currentDate.getMonth() + 1)
					.toString()
					.padStart(2, "0")}-${this.currentDate
					.getDate()
					.toString()
					.padStart(2, "0")}T${paddedEnd}`;

				let payload = {
					_id: this.currentEvent._id,
					ownerID: this.userInfo.ID,
					title: this.form.title,
					desc: this.form.desc,
					start: new Date(newStart).getTime() / 1000,
					end: new Date(newEnd).getTime() / 1000,
				};

				fetch("/backend/API/event/update", {
					method: "POST",
					headers: {
						"Content-Type": "application/json",
						Authorization: "Bearer: " + this.token,
					},
					body: JSON.stringify(payload),
				})
					.then((response) => {
						if (response.status === 200) {
							return response.json();
						} else {
							return "Erro atualizando evento.";
						}
					})
					.then((response) => {
						if (response === "Erro atualizando evento.") {
							console.log(response);
						} else {
							console.log(response);
						}
					});

				// Hide the modal manually
				this.$nextTick(() => {
					this.$bvModal.hide("event-modal");
				});
			}
		},
	},
	computed: {
		...mapGetters(["userInfo"]),
		...mapState(["token", "currentDate", "currentEvent"]),
	},
};
</script>

<style lang="scss" scoped>
</style>