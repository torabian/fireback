import { type DatatableColumn } from "@fireback/ui-core/types/DatatableColumn";
import { UserDto } from "@fireback/manage/sdk/abac/UserDto";
import { type strings as uiStrings } from "@fireback/ui-core/components/strings/translations";
import { type strings } from "./strings/translations";
import { GenderView } from "./GenderView";
import { UserPhotoThumbnail } from "./UserPhotoThumbnail";

export const columns = (
  s: typeof strings,
  uiS: typeof uiStrings,
): DatatableColumn[] => [
  {
    name: UserDto.Fields.uniqueId,
    title: uiS.table.uniqueId,
    width: 100,
  },
  {
    name: UserDto.Fields.firstName,
    filterKey: "first_name",
    title: s.firstName,
    width: 200,
    sortable: true,
    filterable: true,
    getCellValue: (e: UserDto) => e?.firstName,
  },

  {
    filterable: true,
    name: "lastName",
    filterKey: "last_name",
    sortable: true,
    title: s.lastName,
    width: 200,
    getCellValue: (e: UserDto) => e?.lastName,
  },

  {
    name: "birthDate",
    filterKey: "birth_date",
    title: s.birthDate,
    width: 140,
    getCellValue: (e: UserDto) => <>{e?.birthDate}</>,
    filterType: "date",
    filterable: true,
    sortable: true,
  },
  {
    name: "gender",
    title: s.gender,
    width: 50,
    getCellValue: (e: UserDto) => (
      <>
        <GenderView gender={e.gender} />
      </>
    ),
  },

  // {
  //   name: "Image",
  //   title: s.image,
  //   width: 40,
  //   getCellValue: (e: UserDto) => (
  //     <UserPhotoThumbnail
  //       photo={e?.photo}
  //       style={{ width: "20px", height: "20px", objectFit: "cover" }}
  //     />
  //   ),
  // },
  {
    name: UserDto.Fields.primaryAddress.countryCode,
    title: s.countryCode,
    width: 40,
    getCellValue: (e: UserDto) => <>{e.primaryAddress?.countryCode}</>,
  },
  {
    filterable: true,
    name: UserDto.Fields.primaryAddress.addressLine1,
    filterKey: "address_line1",
    title: s.addressLine1,
    width: 180,
    getCellValue: (e: UserDto) => <>{e.primaryAddress?.addressLine1}</>,
  },
  {
    filterable: true,
    name: UserDto.Fields.primaryAddress.addressLine2,
    filterKey: "address_line2",
    title: s.addressLine2,
    width: 180,
    getCellValue: (e: UserDto) => <>{e.primaryAddress?.addressLine2}</>,
  },
  {
    filterable: true,
    name: UserDto.Fields.primaryAddress.city,
    filterKey: "city",
    title: s.city,
    width: 180,
    getCellValue: (e: UserDto) => <>{e.primaryAddress?.city}</>,
  },
  {
    name: UserDto.Fields.primaryAddress.postalCode,
    title: s.postalCode,
    width: 80,
    getCellValue: (e: UserDto) => <>{e.primaryAddress?.postalCode}</>,
  },
  {
    filterable: true,
    name: UserDto.Fields.phoneNumber,
    filterKey: "phone_number",
    title: s.phoneNumber,
    width: 160,
    getCellValue: (e: UserDto) => <>{e?.phoneNumber}</>,
  },
  {
    filterable: true,
    name: UserDto.Fields.jobTitle,
    filterKey: "job_title",
    title: s.jobTitle,
    width: 180,
    getCellValue: (e: UserDto) => <>{e?.jobTitle}</>,
  },
  {
    filterable: true,
    name: UserDto.Fields.company,
    filterKey: "company",
    title: s.company,
    width: 180,
    getCellValue: (e: UserDto) => <>{e?.company}</>,
  },
];
