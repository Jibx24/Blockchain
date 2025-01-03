export interface Appointment {
  index: number; // Block index
  timestamp: string; // เวลาที่บันทึก Block
  previousHash: string; // Hash ของ Block ก่อนหน้า
  hash: string; // Hash ของ Block ปัจจุบัน
  data: {
    patient_id: string; // รหัสคนไข้
    doctor_id: string; // รหัสหมอ
    date: string; // วันที่นัดหมาย
    time: string; // เวลานัดหมาย
    reason: string; // เหตุผลการนัด (ถ้ามี)
    status: string; // สถานะ (Pending, Confirmed, Canceled)
  };
}
