export const dealerInfo = {
  _id: '3',
  dealerName: 'Fremont GMC',
  makeCode: ['Chevrolet', 'Bentley'],
  dealerDoingBusinessAsName: 'FT GMC',
  stateIssuedNumber: '1234_3',
  manufacturerIssuedNumber: '4567_3',
  tenantID: '1',
  website: 'fremont_GMC.cacargroup.com',
  vehicleDamage: [
    {
      vehicleDamageID: '1',
      imageURL: 'https://s3-us-west-1.amazonaws.com/cdms-vehicle-damage-images/Icon-Scratch%403x.png',
      damageType: 'Scratch',
      description: 'Quisque id justo sit amet sapien dignissim vestibulum.',
      priority: 1,
    },
    {
      vehicleDamageID: '2',
      imageURL: 'https://s3-us-west-1.amazonaws.com/cdms-vehicle-damage-images/Icon-Dent%403x.png',
      damageType: 'Dent',
      description: 'Nulla nisl. Nunc nisl. Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa. Donec dapibus. Duis at velit eu est congue elementum. In hac habitasse platea dictumst. Morbi vestibulum, velit id pretium iaculis, diam erat fermentum justo, nec condimentum neque sapien placerat ante.',
      priority: 2,
    },
    {
      vehicleDamageID: '3',
      imageURL: 'https://s3-us-west-1.amazonaws.com/cdms-vehicle-damage-images/Icon-Chipped%403x.png',
      damageType: 'Chip',
      description: 'Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl. Nunc rhoncus dui vel sem. Sed sagittis. Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci.',
      priority: 3,
    },
  ],
  timeZone: 'America/Los_Angeles',
  currency: 'USD',
  dealershipCode: 'ABC123_3',
  dealerGroup: ['1', '2', '3'],
  dealerAddress: [
    {
      dealerAddressID: '1',
      addressType: 'Service',
      streetAddress1: 'Fremont Street No. 999',
      streetAddress2: '',
      city: 'Fremont',
      state: 'CA',
      zipCode: '2365482',
      country: 'US',
      county: '',
      isActive: 'true',
    },
    {
      dealerAddressID: '2',
      addressType: 'Parts',
      streetAddress1: 'New York Street No. 420',
      streetAddress2: '',
      city: 'New York',
      state: 'NY',
      zipCode: '782373',
      country: 'US',
      county: '',
      isActive: 'true',
    },
  ],
  dealerLogos: [
    {
      width: '16',
      height: '16',
      title: 'Icon',
      imageID: 'S3UUIDIcon_1_3',
    },
    {
      width: '48',
      height: '48',
      title: 'Thumb',
      imageID: 'S3UUIDThumb_1_3',
    },
    {
      width: '256',
      height: '256',
      title: 'Original',
      imageID: 'S3UUIDOriginal_1_3',
    },
  ],
  dealerDocumentTemplates: [
    {
      dealerDocumentTemplateID: '1',
      templateName: 'Appointment 1',
      templateType: 'Appointment',
      templateImageID: 'S3ImageID_123',
      isActive: 'true',
    },
    {
      dealerDocumentTemplateID: '2',
      templateName: 'Estimate 1',
      templateType: 'Estimate',
      templateImageID: 'S3ImageID_124',
      isActive: 'true',
    },
    {
      dealerDocumentTemplateID: '3',
      templateName: 'Repair Order 1',
      templateType: 'Repair Order',
      templateImageID: 'S3ImageID_125',
      isActive: 'false',
    },
  ],
  dealerOperationSchedule: [
    {
      dealerOperationScheduleID: '1',
      dealerOperationType: 'Sales',
      sundayOpenHour: '7:00 AM',
      sundayCloseHour: '7:00 PM',
      mondayOpenHour: '7:00 AM',
      mondayCloseHour: '7:00 PM',
      tuesdayOpenHour: '7:00 AM',
      tuesdayCloseHour: '7:00 PM',
      wednesdayOpenHour: '7:00 AM',
      wednesdayCloseHour: '7:00 PM',
      thursdayOpenHour: '7:00 AM',
      thursdayCloseHour: '7:00 PM',
      fridayOpenHour: '7:00 AM',
      fridayCloseHour: '7:00 PM',
      saturdayOpenHour: '7:00 AM',
      saturdayCloseHour: '7:00 PM',
    },
    {
      dealerOperationScheduleID: '2',
      dealerOperationType: 'Parts',
      sundayOpenHour: '8:00 AM',
      sundayCloseHour: '8:00 PM',
      mondayOpenHour: '8:00 AM',
      mondayCloseHour: '8:00 PM',
      tuesdayOpenHour: '8:00 AM',
      tuesdayCloseHour: '8:00 PM',
      wednesdayOpenHour: '8:00 AM',
      wednesdayCloseHour: '8:00 PM',
      thursdayOpenHour: '8:00 AM',
      thursdayCloseHour: '8:00 PM',
      fridayOpenHour: '8:00 AM',
      fridayCloseHour: '8:00 PM',
      saturdayOpenHour: '8:00 AM',
      saturdayCloseHour: '8:00 PM',
    },
  ],
  dealerContact: ['1', '2'],
  isActive: true,
  lastUpdatedByUser: 'Qasim',
  lastUpdatedByDisplayName: 'DealerService',
  lastUpdatedDateTime: '',
  documentVersion: '1.0',
};